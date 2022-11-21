package main

import "fmt"

func main() {
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
