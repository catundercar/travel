package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	MAXLENGTH = 10000
)

var (
	test1 = make([]int, MAXLENGTH)
	test2 = make([]int, MAXLENGTH)
	test3 = make([]int, MAXLENGTH)
	test4 = make([]int, MAXLENGTH)
)

func init() {
	for i := 0; i < MAXLENGTH; i++ {
		rand.Seed(time.Now().Unix())
		randNum := rand.Intn(MAXLENGTH)
		test1[i] = randNum
		test2[i] = randNum
		test3[i] = randNum
		test4[i] = randNum
	}
	fmt.Println("Data Length: ", MAXLENGTH)
	//fmt.Println("example: ", test1)
}

func TestBubbleSort(t *testing.T) {
	a := []int{3, 5, 7, 3, 1, 4, 5, 9, 12}
	BubbleSort(a)
	fmt.Println(a)
}

func TestInsertionSort(t *testing.T) {
	a := []int{3, 5, 7, 3, 1, 4, 5, 9, 12}
	InsertionSort(a)
	fmt.Println(a)
}

func TestMergeSort(t *testing.T) {
	a := []int{3, 5, 7, 3, 1, 4, 5, 9, 12}
	MergeSort(a)
	fmt.Println(a)
}

func TestSelectSort(t *testing.T) {
	a := []int{3, 5, 7, 3, 1, 4, 5, 9, 12}
	SelectSort(a)
	fmt.Println(a)
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(test1)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(test2)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(test3)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectSort(test4)
	}
}
