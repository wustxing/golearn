package heapdefine

import (
	"fmt"
	"testing"
)

func Test_Heap(t *testing.T) {
	var h = Heap{20, 7, 3, 10, 15}

	h.Init()
	fmt.Println(h)

	h.Push(6)
	fmt.Println(6)

	x, ok := h.Remove(5)
	fmt.Println(x, ok, h)

	z := h.Pop()
	fmt.Println(z, h)
}
