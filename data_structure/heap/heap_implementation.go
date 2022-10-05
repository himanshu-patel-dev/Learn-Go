package main

import (
	"fmt"
)

// struct that holds the data
type MaxHeap struct {
	data []int
}

// add element to heap
func (h *MaxHeap) Insert(key int) {
	h.data = append(h.data, key)
	h.maxHeapifyUp(len(h.data) - 1) // index which we want to heapify up
}

// remove max element from top of heap
func (h *MaxHeap) Pop() int {
	if len(h.data) == 0 {
		return -1
	}

	mx_ele := h.data[0]
	l := len(h.data) - 1
	h.data[0] = h.data[l] // remove first ele by over writing from last element
	h.data = h.data[:l]   // remove last element by dec size of array
	h.maxHeapifyDown(0)   // index which we want to heapify down
	return mx_ele
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	left, right := left(index), right(index)
	l := len(h.data)

	if left < l && right < l {
		left_data, right_data, curr_data := h.data[left], h.data[right], h.data[index]
		if left_data > curr_data && left_data >= right_data {
			h.swap(left, index)
			h.maxHeapifyDown(left)
		} else if right_data > curr_data && right_data >= left_data {
			h.swap(right, index)
			h.maxHeapifyDown(right)
		}
	}
	// if only left child is available and is more in value than parent
	if left < len(h.data) && h.data[left] > h.data[index] {
		h.swap(left, index)
		h.maxHeapifyDown(left)
	}
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for index > 0 && h.data[parent(index)] < h.data[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return index*2 + 1
}

func right(index int) int {
	return index*2 + 2
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	maxHeap := []int{10, 20, 30, 15, 17}
	for _, v := range maxHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	fmt.Println("----------------------------")
	l := len(m.data)
	for i := 0; i < l; i++ {
		fmt.Println(m.data)
		fmt.Println(m.Pop())
	}
}
