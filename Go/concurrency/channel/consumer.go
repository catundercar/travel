package main

import (
	"fmt"
	"sync"
)

func worker(ch chan int, done chan bool) {
	mux := sync.Mutex{}
	a := 0
	routines := 2
	wg := sync.WaitGroup{}
	for i := 0; i < routines; i++ {
		wg.Add(1)
		go func(a int) {
			defer wg.Done()

			for {
				select {
				case x, ok := <-ch:
					if !ok {
						fmt.Println("xxx")
						return
					}

					mux.Lock()
					a += 1
					mux.Unlock()
					fmt.Println(x)
				}
			}
		}(a)
	}
	wg.Wait()
	fmt.Println(a)
	done <- false
}

func producer(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func master() {

	ch := make(chan int, 10)
	done := make(chan bool)
	go producer(ch)
	go worker(ch, done)
	fmt.Println(<-done)
	fmt.Println("work done")

	//	time.Sleep(5 * time.Second)
	//	for {
	//		if <-done {
	//			return
	//		}
	//	}
	//	select {
	//	case x, ok := <-done:
	//		if ok {
	//			if x {
	//				fmt.Println("done")
	//			}
	//			return
	//		}
	//	}
}
