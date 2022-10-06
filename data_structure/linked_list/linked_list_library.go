package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Initialize an empty list
	l := list.New()
	fmt.Println(l)

	fmt.Println("\nEmpty Elements")
	fmt.Println(l.Front())
	fmt.Println(l.Back())
	fmt.Println(l.Len())

	// add item to front
	l.PushFront(10)
	fmt.Println("\nFrom Front")
	fmt.Println(l.Front())
	fmt.Println(l.Front().Value, l.Front().Next(), l.Front().Prev())

	fmt.Println("\nAfter pushback")
	l.PushBack(15)
	l.PushBack(20)
	fmt.Println(l)

	fmt.Println("\nIterate over linkedlist")
	e := l.Front()
	for e != nil {
		fmt.Println(e.Value)
		e = e.Next()
	}

	fmt.Println() // OR
	for p := l.Front(); p != nil; p = p.Next() {
		fmt.Println(p.Value)
	}
}
