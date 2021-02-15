package main

func hello() {
	// will print "Hello, World" at some point in the future (perhaps after hello has returned)
	a := "Hello World"
	go func() {
		print(a)
	}()
}
