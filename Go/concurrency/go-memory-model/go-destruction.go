package main

func destruction() {
	var a string
	go func() {
		// can not guaranteed to be observed by any other goroutine.
		a = "Hello World"
	}()
	print(a)
}

