package main

import (
	"fmt"
	"runtime"
)

type taskFunc func()

// DoTaskConcurrent can do tasks concurrently.The concurrentNum is used to set the number of
// goroutines.
func DoTaskConcurrent(concurrentNum int, fns chan taskFunc) {
	fmt.Println(runtime.GOMAXPROCS(4))
	taskNum := len(fns)
	fmt.Println(taskNum)
	taskChan := make(chan struct{}, taskNum)
	c := make(chan struct{}, concurrentNum)
	defer close(taskChan)
	defer close(c)

	for i := 0; i < taskNum; i++ {
		c <- struct{}{}
		go func() {
			fn := <-fns
			fn()
			<-c
			taskChan <- struct{}{}
		}()
	}

	for i := taskNum; i > 0; i-- {
		<-taskChan
	}
}
