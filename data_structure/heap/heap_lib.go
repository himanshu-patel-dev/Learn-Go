package main

import (
	"container/heap"
	"fmt"
)

// declare a heap type
type IntHeap []int

// implement 3 mandatory method on the heap data structure as per heap lib interface
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Println("Heap: ", (*h))
	for h.Len() > 0 {
		fmt.Printf("> %d \n", heap.Pop(h))
	}
	fmt.Println()
}

// Heap:  [1 2 5 3]
// > 1
// > 2
// > 3
// > 5
