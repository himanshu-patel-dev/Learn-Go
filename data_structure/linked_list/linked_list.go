package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

// given the pointer to new node we can add that in front of
// existing LL
func (l *linkedList) prepend(n *node) {
	n.next = l.head
	l.head = n
	l.length++
}

func main() {
	ll := linkedList{}
	node1 := node{data: 1}
	node2 := node{data: 2}
	node3 := node{data: 3}
	node4 := node{data: 4}
	ll.prepend(&node1)
	ll.prepend(&node2)
	ll.prepend(&node3)
	ll.prepend(&node4)

	temp := ll.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

/*
4
3
2
1
*/
