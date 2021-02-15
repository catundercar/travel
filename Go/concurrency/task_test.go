package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestDoTaskConcurrent(t *testing.T) {
	mocks := mockTask()
	DoTaskConcurrent(10, mocks)
}

func TestDoTaskSync(t *testing.T) {
	mocks := mockTask()
	fmt.Println(len(mocks))
	for i := 0; i < len(mocks); i++ {
		fn := <-mocks
		fn()
	}
}

func mockTask() chan taskFunc {
	mocks := make(chan taskFunc, 100)
	for i := 0; i < 100; i++ {
		a := rand.Intn(100)
		b := rand.Intn(100)
		mocks <- mockAnd(a, b)
	}
	return mocks
}

func mockAnd(a, b int) taskFunc {
	return func() {
		num1 := a
		num2 := b
		fmt.Printf("a + b = %d\n", num1+num2)
		time.Sleep(time.Millisecond * 1000)
	}
}
