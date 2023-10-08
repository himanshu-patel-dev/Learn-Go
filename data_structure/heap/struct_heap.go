package main

import (
	"container/heap"
	"fmt"
)

type MinNode struct {
	priority int
	data     []int
}

type MinHeap []*MinNode

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MinHeap) Less(i, j int) bool { return m[i].priority < m[j].priority }

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(*MinNode))
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	min := old[n-1]
	*m = old[0 : n-1]
	return min
}

func main() {
	n1 := &MinNode{priority: 4, data: []int{4}}
	n2 := &MinNode{priority: 2, data: []int{2}}
	n3 := &MinNode{priority: 1, data: []int{1}}
	n4 := &MinNode{priority: 0, data: []int{0}}
	pq := &MinHeap{n1, n2, n3}
	// be carefull here, it is heap.Init not minHeap.Init, use the heap package not struct
	heap.Init(pq)
	// be carefull here, it is heap.Push not minHeap.Push, use the heap package not struct
	heap.Push(pq, n4)
	printAll(pq)
	fmt.Println("-----------")
	for pq.Len() > 0 {
		// be carefull here, it is heap.Pop not minHeap.Pop, use the heap package not struct
		node := (heap.Pop(pq)).(*MinNode)
		fmt.Println(node.priority)
	}
}

func printAll(minHeap *MinHeap) {
	for i := range *minHeap {
		fmt.Printf("%v\t", (*minHeap)[i])
	}
	fmt.Println()
}
