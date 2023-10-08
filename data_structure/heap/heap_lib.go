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
	heap.Init(h)	// this makes an array a valid min heap
	// after init every push and pop will retain a valid heap
	heap.Push(h, 6)
	heap.Push(h, 3)
	heap.Push(h, 4)
	heap.Push(h, 0)

	// fmt.Println("Heap: ", (*h))
	for h.Len() > 0 {
		fmt.Println("Heap: ", (*h))

		fmt.Printf("> %d \n", heap.Pop(h))
	}
	fmt.Println()
}

// Heap:  [0 2 1 6 3 5 4]
// > 0 
// Heap:  [1 2 4 6 3 5]
// > 1 
// Heap:  [2 3 4 6 5]
// > 2 
// Heap:  [3 5 4 6]
// > 3 
// Heap:  [4 5 6]
// > 4 
// Heap:  [5 6]
// > 5 
// Heap:  [6]
// > 6 