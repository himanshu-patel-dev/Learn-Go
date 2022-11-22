package main

import "fmt"

// an int parameter, so arguments will be passed to it by value
func byValue(i int) int {
	i = 0
	return i
}

// *int parameter, meaning that it takes an int pointer.
func byPointer(i *int) {
	*i = 0 // dereference and assign value
}

func main() {
	i := 5

	fmt.Println(byValue(i)) // 0
	fmt.Println(i)          // 5
	byPointer(&i)
	fmt.Println(i) // 0
}
