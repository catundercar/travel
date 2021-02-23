package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// 设置计算器为2
	wg.Add(2)

	go func() {
		// 每个Done计算器减一
		defer wg.Done()
		for i := 1; i < 100; i++ {
			fmt.Println("A: ", i)
			//time.Sleep(10 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			fmt.Println("B: ", i)
			//time.Sleep(100 * time.Millisecond)
		}
	}()

	// 阻塞， 等待计算器为0
	wg.Wait()
	fmt.Println(runtime.NumCPU())
}
