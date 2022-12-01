package main

import (
	"fmt"
	"sort"
)

// byLength type that is just an alias for the builtin []string type
type byLength []string

// defining 3 function on byLength interface
// which are needed for sort to work

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits) // [kiwi peach banana]
}
