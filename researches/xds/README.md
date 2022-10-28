# xds（xmap）库调研

[TOC]

## 目的

为了提高sync.map的性能，找了一个号称比sync.map性能有 2 倍的提升。Github地址为：

[heiyeluren/xds: A third-party extensible collection of high-performance data structures and data types in Go. 第三方可扩展的Go语言中高性能数据结构和数据类型合集 (github.com)](https://github.com/heiyeluren/xds)



## 性能对比

下面从QPS、内存占用两个方面做性能比较。map容量为 800,000.

### xmap

* 内存

  ```go
    keys := make([]string, 8000000)
  	for i := 0; i < 8000000; i++ {
  		keys[i] = strconv.Itoa(rand.Int() % 800000)
  	}
  	for _, key := range keys {
  		if err := chm.Put([]byte(key), []byte(key+"xx")); err != nil {
  			panic(err)
  		}
  	}
  ```

  ![xmap-mem](png\xmap-mem.png)

  内存常驻在1.5g 左右。

* Benchmark(SetParallelism(100))

  ```go
  BenchmarkCHM_Concurrent_Get-32           5417234               260.7 ns/op
  ```

* ab测试结果

  ab测试设置一共请求Get（Read） 100,000次。

  | 并发级别 | 总耗时        | 平均每秒处理请求数 | 平均请求耗时（一次并发） | 平均请求耗时（一次请求） |
    | -------- | ------------- | ------------------ | ------------------------ | ------------------------ |
  | 100      | 2.903 seconds | 34451.12           | 2.903ms                  | 0.029ms                  |
  | 1000     | 2.820 seconds | 35455.50           | 28.204ms                 | 0.028ms                  |
  | 10000    | 2.992 seconds | 33427.64           | 299.154ms                | 0.030ms                  |



### sync.Map

* 内存

  ```go
  keys := make([]string, 8000000)
  	for i := 0; i < 8000000; i++ {
  		keys[i] = strconv.Itoa(rand.Int() % 800000)
  	}
  	for _, key := range keys {
  		f.Store(key, key+"xx")
  	}
  ```

  ![sync-map-mem](png\sync-map-mem.png)

  内存常驻在 495.8m

* Benchmark(SetParallelism(100))

  ```go
  BenchmarkSyncMap_Get-32          4061788               275.6 ns/op
  ```

* ab测试结果

  ab测试设置一共请求Get（Read） 100,000次。

| 并发级别 | 总耗时        | 平均每秒处理请求数 | 平均请求耗时（一次并发） | 平均请求耗时（一次请求） |
| -------- | ------------- | ------------------ | ------------------------ | ------------------------ |
| 100      | 2.498 seconds | 40029.54           | 2.498ms                  | 0.025ms                  |
| 1000     | 2.963 seconds | 33750.34           | 29.629ms                 | 0.030ms                  |
| 10000    | 2.999 seconds | 33349.77           | 299.852ms                | 0.030ms                  |

## 结论

从内存上看，`xmap`比`sync.Map`内存占用多了三倍，从QPS 来说，在并发级别低时，`sync.Map`略胜一筹，并发级别为1000或者10000时，二者相差无几。但是从进程内访问（函数调用）来看，`xmap`比`sync.Map`更快，1s可以多读取一百多万次。
