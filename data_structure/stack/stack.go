package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() int {
	l := len(s.items)
	if l == 0 {
		return -1
	}
	result := s.items[l-1]
	s.items = s.items[:l-1]
	return result
}

func main() {
	myStack := &Stack{}
	myStack.Push(10)
	myStack.Push(20)
	myStack.Push(30)
	fmt.Println(myStack)

	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack)
}

/*
&{[10 20 30]}
30
20
10
&{[]}
*/
