package main

var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	// A send on a channel happens before the corresponding receive from that channel completes.
	// c <- 0
	// The closing of a channel happens before a receive that returns a zero value because the channel has closed.
	close(c)
}

func main() {
	go f()
	<-c
	print(a)
}
