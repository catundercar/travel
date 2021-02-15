package main

import (
	"sync"
)

var d string
var once sync.Once

func setup() {
	d = "hello, world"
	print("init d")
}

func doprint() {
	once.Do(setup)
	print(d)
}

func twoprint() {
	// print once "init d", twice "hello world"
	// A single call of f() from once.Do(f) happens (returns) before any call of once.Do(f) returns.
	doprint()
	doprint()
}

func main()  {
	twoprint()
}