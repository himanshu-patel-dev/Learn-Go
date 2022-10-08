package main

import "fmt"

const AlphabetSize = 26

// Node is node of a trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represet a trie and has pointer to root node
type Trie struct {
	root *Node
}

// InitTrie Initialize a trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take a word and insert that into a trie
func (t *Trie) Insert(word string) {
	wordLen := len(word)
	currentNode := t.root
	for i := 0; i < wordLen; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(word string) bool {
	wordLen := len(word)
	currentNode := t.root
	for i := 0; i < wordLen; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd
}

func main() {
	myTrie := InitTrie()
	// fmt.Println(myTrie.root)
	myTrie.Insert("dog")
	fmt.Println(myTrie.Search("dogs"))
	fmt.Println(myTrie.Search("dog"))
}
