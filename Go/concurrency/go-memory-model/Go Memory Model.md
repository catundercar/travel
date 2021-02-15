# Go Memory Model

## 介绍

如何保证在一个goruntine中看到在另一个goroutine修改的变量的值

如果程序中修改数据时有其他goroutine同时读取，那么必须将读取串行化。为了串行化访问，请使用channel或其他同步原语，例如sync和sync/atomic来保护数据。

## Happens Before

在一个gouroutine中，读和写一定是按照程序中的顺序执行的。即编译器和处理器只有在不会改变这个goroutine的行为时才可能修改读和写的执行顺序。**由于重排，不同的goroutine可能会看到不同的执行顺序。**例如，一个goroutine执行`a = 1;b = 2;`，另一个goroutine可能看到`b`在`a`之前更新。

为了说明读和写的必要条件，我们定义了`先行发生（Happens Before）`--Go程序中执行内存操作的偏序。如果事件`e1`发生在`e2`前，我们可以说`e2`发生在`e1`后。如果`e1`不发生在`e2`前也不发生在`e2`后，我们就说`e1`和`e2`是并发的。

### 内存重排

- CPU重排

[MESI](https://zh.wikipedia.org/wiki/MESI%E5%8D%8F%E8%AE%AE)

- 编译器重排

## 同步

### 初始化

程序的初始化在单独的goroutine中进行，但这个goroutine可能会创建出并发执行的其他goroutine。

*If a package `p` imports package `q`, the completion of `q`'s `init` functions happens before the start of any of `p`'s.*

*The start of the function `main.main` happens after all `init` functions have finished.*

***如果包p引入（import）包q，那么q的init函数的结束先行发生于p的所有init函数开始*** 。

***main.main函数的开始发生在所有init函数结束之后***

### 创建goroutine

The go statement that starts a new goroutine happens before the goroutine's execution begins.

启动一个goroutine的语句先行发生于goroutine的执行。

```go
func main() {
	// will print "Hello, World" at some point in the future (perhaps after main has returned)
	a := "Hello World"
	go func() {
		print(a)
	}()
}
```

will print "Hello, World" at some point in the future (perhaps after *main* has returned)

### 销毁goroutine

gouroutine的退出并不会保证先行发生于程序的任何事件。

```go
var a string

func hello() {
	go func() { a = "hello" }()
	print(a)
}
```

对变量`a` 的赋值之后没有任何同步事件，所以不能保证变量`a` 对其他goroutine可见。实际上，一个激进的编译器可能会删掉整个`go` 语句。

### Channel

channel通信是goroutine同步的主要方法。每一个在特定channel的发送操作都会匹配到通常在另一个goroutine执行的接收操作。

- *A send on a channel happens before the corresponding receive from that channel completes.*

*对channel的发送操作先行发生于**对应的**接收操作完成。*

```go
var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	c <- 0
}

func main() {
	go f()
	<-c
	print(a)
}
```

这段程序保证了变量`a` 的赋值的可见性。`a`的赋值先行发生于`c <- 0`，`c <- 0`先行发生于`<-c`，`<-c`先行发生于`print(a)`,所以`a`的赋值先行发生于`print(a)` 。

- *The closing of a channel happens before a receive that returns a ze ro value because the channel is closed.*

*一个`channel`的关闭先行发生于它接收到一个零值，因为它已经关闭了。*

在上面的例子中，将`c <- 0`替换为`close(c)`还会产生同样的结果。

- *A receive from an unbuffered channel happens before the send on that channel completes.*

*无缓冲`channel`的接收先行发生于发送完成.*

```go
var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	<-c
}

func main() {
	go f()
	c <- 0
	print(a)
}
```

- *The kth receive on a channel with capacity C happens before the k+Cth send from that channel completes.*

*在容量为C的channel上的第k个接收先行发生于从这个channel上的第k+C次发送完成。*

这条规则将前面的规则推广到了带缓冲的channel上。可以通过带缓冲的channel来实现计数信号量：channel中的元素数量对应着活动的数量，channel的容量表示同时活动的最大数量，发送元素获取信号量，接收元素释放信号量，这是限制并发的通常用法。

下面程序为`work`中的每一项开启一个goroutine，但这些goroutine通过有限制的channel来确保最多同时执行三个工作函数（w）。

```go
var limit = make(chan int, 3)

func main() {
	for _, w := range work {
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	select{}
}
```

### 锁

sync包实现了两个锁的数据类型`sync.Mutex`和`sync.RWMutex`。

- *For any sync.Mutex or sync.RWMutex variable l and n < m, call n of l.Unlock() happens before call m of l.Lock() returns.*

*对任意的sync.Mutex或sync.RWMutex变量l和n < m，n次调用l.Unlock()先行发生于m次l.Lock()返回。*

```go
var l sync.Mutex
var a string

func f() {
	a = "hello, world"
	l.Unlock()
}

func main() {
	l.Lock()
	go f()
	l.Lock()
	print(a)
}
```

能保证打印出"hello, world"。第一次调用l.Unlock()（在f()中）先行发生于main中的第二次l.Lock()返回, 先行发生于print。

- *For any call to l.RLock on a sync.RWMutex variable l, there is an n such that the l.RLock happens (returns) after call n to l.Unlock and the matching l.RUnlock happens before call n+1 to l.Lock.*

*对于sync.RWMutex变量l，任意的函数调用l.RLock满足第n次l.RLock**后发生于**第n次调用l.Unlock，对应的l.RUnlock先行发生于第n+1次调用l.Lock。*

### Once

sync包的Once为多个goroutine提供了安全的初始化机制。能在多个线程中执行once.Do(f)，但只有一个f()会执行，其他调用会一直阻塞直到f()返回。

*A single call of f() from [once.Do](http://once.do/)(f) happens (returns) before any call of [once.Do](http://once.do/)(f) returns.*

```go
var a string
var once sync.Once

func setup() {
	a = "hello, world"
}

func doprint() {
	once.Do(setup)
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}
```

调用`twoprint`会打印*"hello, world"*两次。`setup`只在第一次`doprint`时执行。

### 错误的同步方法

注意，读操作***r***可能会看到并发的写操作***w***。即使这样也不能表明***r***之后的读能看到***w***之前的写。

如下程序：

```go
var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	print(b)
	print(a)
}

func main() {
	go f()
	g()
}
```

`g`可能先打印出2然后是0。

![](D:\go\src\CS-Study\pics\golang\Go-Memory-Model.png)