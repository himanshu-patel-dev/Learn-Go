package main

import "fmt"

// ArraySize is the size of hash table array
const ArraySize = 7

// HashTable will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

// Insert will take in a key and add it to hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search will search for the given key in hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will delete the key from hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// bucket is linked list in each slot
type bucket struct {
	head *bucketNode
}

// insert put the given key in linked list / bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		// append the newNode to start of linked list
		newNode.next = b.head
		b.head = newNode
	}
}

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete will take a string/key and remove it from bucket
func (b *bucket) delete(k string) {
	// if delete key is the head then reset the head to next node
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == k {
			prevNode.next = prevNode.next.next
			return
		}
		prevNode = prevNode.next
	}
}

// bucketNode is linked list node that hold a key
type bucketNode struct {
	key  string
	next *bucketNode
}

// hash takes a string an return a hash (int) which is
// basically a sum of ASCII value of all char in string
// modulo ArraySize
func hash(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of hash table
func Init() *HashTable {
	ht := &HashTable{}
	for i := range ht.array {
		ht.array[i] = &bucket{}
	}
	return ht
}

func main() {
	myHash := Init()
	list := []string{
		"SPIDERMAN",
		"PERMAN",
		"SUPERMAN",
	}
	for _, hero := range list {
		myHash.Insert(hero)
	}
	fmt.Println(myHash.Search("SPIDERMAN")) // true
	fmt.Println(myHash.Search("SHAKTIMAN")) // false
	myHash.Delete("SPIDERMAN")
	fmt.Println(myHash.Search("SPIDERMAN")) // false
}

/*
true
false
false
*/
