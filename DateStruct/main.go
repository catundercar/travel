package main

import (
	"CS-Study/DateStruct/sorts"
	"fmt"
)

func main() {
	a := []int{1, 5, 4, 2, 6, 9, 3, 4, 7}
	fmt.Println(a)
	// sorts_v1.Bubblesort(a)
	// sorts_v1.InsertionSort(a)
	sorts.MergeSort(a, len(a))
	fmt.Println(a)
}
