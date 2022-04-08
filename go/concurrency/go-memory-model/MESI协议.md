# MESI协议

MESI协议是基于失效机制的缓存一致性协议，并且是支持*回写（write-back）*缓存的最常用协议。

也称为*Illinois Protocol。*与*写通过（write through）*相比，回写缓存能节约大量带宽。总是有“脏“（dirty）状态表示缓存中的数据与主存不同。MESI协议要求在缓存不命中（miss）且数据块在另一个缓存时，允许缓存到缓存的数据复制。与MSI协议相比，MESI协议减少了主存的事务数量，这极大的改善了性能。

## 状态

缓存行有4种状态：

- M（modified）已修改状态

缓存行是脏的（dirty），与主存的值不同。如果别的CPU内核要读主存这块数据，该缓存行必须回写到主存，状态变为共享（S）。

- E（exclusive）独占的

缓存行只在当前缓存中，但是是干净的（clean）—缓存数据同于主存数据。当别的缓存读取它时，状态变为共享；当前写数据时，变为已修改状态。

- S（shared）共享的

缓存行也存在其他缓存中且是干净的。缓存行可以在任意时刻抛弃。

- I（invalid）失效的

缓存行是无效的。

对于任意一对缓存，对应缓存行的相容关系

|      | M                  | E                  | S                  | I                  |
| ---- | ------------------ | ------------------ | ------------------ | ------------------ |
| M    | :x:                | :x:                | :x:                | :heavy_check_mark: |
| E    | :x:                | :x:                | :x:                | :heavy_check_mark: |
| S    | :x:                | :x:                | :heavy_check_mark: | :heavy_check_mark: |
| I    | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |

当缓存行标记为失效时，在其他缓存中的数据副本被标记为I（无效）。

## 操作

有限状态自动机的状态转换结束两种场景：缓存所在处理器的读写；其他处理器的读写。总线请求被称为总线窥探器监视。

处理器对缓存的请求：

1. PrRd：处理器请求**读**一个缓存块
2. PrWr：处理器请求**写**一个缓存块

总线对缓存的请求：

1. BusRd：监听到请求，表示当前有其它处理器正在发起对某个 Cache block 的读请求
2. BusRdx：监听到请求，表示当前有其它处理器正在发起对某个其未拥有的 Cache block 的写请求
3. BusUpgr： 监听到请求，表示有另一个处理器正在发起对某 Cache block 的写请求，该处理器已经持有此 Cache block 块
4. Flush： 监听到请求，表示整个 cache 块已被另一个处理器写入到主存中
5. FlushOpt： 监听到请求，表示一个完整的 cache 块已经被发送到总线，以提供给另一个处理器使用(Cache 到 Cache 传数)

Cache 到 Cache 的传送可以降低 read miss 导致的延迟，如果不这样做，需要先将该 block 写回到主存，再读取，延迟会大大增加，在基于总线的系统中，这个结论是正确的。但在多核架构中，coherence 是在 L2 caches 这一级保证的，从 L3 中取可能要比从另一个 L2 中取还要快。

mesi 协议解决了多核环境下，内存多层级带来的问题。使得 cache 层对于 CPU 来说可以认为是透明的，不存在的。单一地址的变量的写入，可以以线性的逻辑进行理解。

但 mesi 协议有两个问题没有解决，一种是 RMW 操作，或者叫 CAS；一种是 ADD 操作。因为这两种操作都需要先读到原始值，进行修改，然后再写回到内存中。

同时，在 CPU 架构中我们看到 CPU 除了 cache 这一层之外，还存在 store buffer，而 store buffer 导致的内存乱序问题，mesi 协议是解决不了的，这是 memory consistency 范畴讨论的问题。



## **store buffer 和 invalidate queue**

> Store Buffer：

> 当写入到一行 invalidate 状态的 cache line 时便会使用到 store buffer。写如果要继续执行，CPU 需要先发出一条 read-invalid 消息(因为需要确保所有其它缓存了当前内存地址的 CPU 的 cache line 都被 invalidate 掉)，然后将写推入到 store buffer 中，当最终 cache line 达到当前 CPU 时再执行这个写操作。

> CPU 存在 store buffer 的直接影响是，当 CPU 提交一个写操作时，这个写操作不会立即写入到 cache 中。因而，无论什么时候 CPU 需要从 cache line 中读取，都需要先扫描它自己的 store buffer 来确认是否存在相同的 line，因为有可能当前 CPU 在这次操作之前曾经写入过 cache，但该数据还没有被刷入过 cache(之前的写操作还在 store buffer 中等待)。需要注意的是，虽然 CPU 可以读取其之前写入到 store buffer 中的值，但其它 CPU 并不能在该 CPU 将 store buffer 中的内容 flush 到 cache 之前看到这些值。即 store buffer 是不能跨核心访问的，CPU 核心看不到其它核心的 store buffer。

> Invalidate Queues：

> 为了处理 invalidation 消息，CPU 实现了 invalidate queue，借以处理新达到的 invalidate 请求，在这些请求到达时，可以马上进行响应，但可以不马上处理。取而代之的，invalidation 消息只是会被推进一个 invalidation 队列，并在之后尽快处理(但不是马上)。因此，CPU 可能并不知道在它 cache 里的某个 cache line 是 invalid 状态的，因为 invalidation 队列包含有收到但还没有处理的 invalidation 消息，这是因为 CPU 和 invalidation 队列从物理上来讲是位于 cache 的两侧的。

> 从结果上来讲，memory barrier 是必须的。一个 store barrier 会把 store buffer flush 掉，确保所有的写操作都被应用到 CPU 的 cache。一个 read barrier 会把 invalidation queue flush 掉，也就确保了其它 CPU 的写入对执行 flush 操作的当前这个 CPU 可见。再进一步，MMU 没有办法扫描 store buffer，会导致类似的问题。这种效果对于单线程处理器来说已经是会发生的了。

因为前面提到的 store buffer 的存在，会导致多核心运行用户代码时，读和写以非程序顺序的顺序完成。