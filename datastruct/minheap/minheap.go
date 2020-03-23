package minheap

import "math"

type MinHeap struct {
	elements []int
}

func NewMinHeap() *MinHeap {
	h := &MinHeap{elements: []int{math.MinInt64}}
	return h
}

func (h *MinHeap) Insert(v int) {
	h.elements = append(h.elements, v)
	i := len(h.elements) - 1
	for ; h.elements[i/2] > v; i /= 2 {
		h.elements[i] = h.elements[i/2]
	}
	h.elements[i] = v
}
