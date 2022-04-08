package main

var unbufferedChan = make(chan struct{})

var b string

func fn() {
	b = "Hello World"
	// a receive from an unbuffered channel happens before the send on that channel completes.
	<- unbufferedChan
}

func main() {
	go fn()
	unbufferedChan <- struct{}{}
	print(b)
}
