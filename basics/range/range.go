package main

import "fmt"

// The type of elements and length are both part of the array’s type.
// By default an array is initialized with zero-valued of it's type

func main() {
	// create an empty array with all values init with zero value
	var a [5]int
	fmt.Println("emp:", a)

	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index:", i, "value: ", value)
	}
	// index: 0 value:  0
	// index: 1 value:  1
	// index: 2 value:  -1
	// index: 3 value:  2
	// index: 4 value:  -2

	fmt.Println("--------------")
	for _, value := range anArray {
		fmt.Println("value: ", value)
	}
	// value:  0
	// value:  1
	// value:  -1
	// value:  2
	// value:  -2

	fmt.Println("--------------")
	for index := range anArray {
		fmt.Println("index: ", index)
	}
	// index:  0
	// index:  1
	// index:  2
	// index:  3
	// index:  4
}
