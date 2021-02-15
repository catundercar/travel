package main

import "sync"

var l sync.Mutex
var lc string

func fnl() {
	lc = "hello, world"
	// For any sync.Mutex or sync.RWMutex variable l and n < m, call n of l.Unlock() happens before call m of l.Lock() returns.
	l.Unlock()
}

func main() {
	l.Lock()
	go fnl()
	l.Lock()
	print(lc)
}
