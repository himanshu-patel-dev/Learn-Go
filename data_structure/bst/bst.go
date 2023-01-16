package main

import "fmt"

// Node to BST
type Node struct {
	data  int
	left  *Node
	right *Node
}

// Insert node in BST
func (n *Node) Insert(value int) *Node {
	if n == nil {
		return &Node{data: value}
	}

	if n.data > value {
		n.left = n.left.Insert(value)
	} else if n.data < value {
		n.right = n.right.Insert(value)
	}
	return n
}

// Search for value in BST
func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}
	if n.data == value {
		return true
	} else if n.data < value {
		return n.right.Search(value)
	} else {
		return n.left.Search(value)
	}
}

func main() {
	tree := &Node{data: 100}
	tree.Insert(50)
	tree.Insert(150)
	fmt.Println(tree.left.data, tree.data, tree.right.data)
	fmt.Println(tree.Search(150))
	fmt.Println(tree.Search(151))
}

/*
50 100 150
true
false
*/
