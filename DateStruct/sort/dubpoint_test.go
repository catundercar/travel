package sort

import (
	"fmt"
	"testing"
)

func TestMoveZero(t *testing.T) {
	// t.Fatal("not implemented")
	a := []int{0, 1, 2, 3, 4, 0, 5, 3, 0}
	fmt.Println(a)
	movezore(a)
	fmt.Println(a)
}
