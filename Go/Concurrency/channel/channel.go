package main

import "fmt"

func deadlock1() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch2 <- "ch2 value"
		ch1 <- "ch1 value"
	}()

	fmt.Println(<-ch1)
}

func main() {
	//	deadlock1()
	master()
}
