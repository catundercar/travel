package sort

import (
	"fmt"
	"testing"
)

func TestMoveZero(t *testing.T) {
	// t.Fatal("not implemented")
	a := []int{1, 1, 1, 1, 0, 0, 0, 0, 1}
	fmt.Println(a)
	movezore(a)
	fmt.Println(a)
}
