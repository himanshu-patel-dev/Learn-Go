package main

import "fmt"

func printSlice(slice []int) {
	fmt.Println("----------------")
	fmt.Println(slice)
	fmt.Printf("%T\n", slice)
	fmt.Println(len(slice), cap(slice))
	fmt.Println("----------------")
}

func main() {
	newSlice := []int{1, 2, 3, 4, 5}
	printSlice(newSlice)

	// keep allocated memory - slice the array to zero length
	// but the capacity and type remians the same
	newSlice = newSlice[:0]
	printSlice(newSlice)

	// if slice is extended the original data reappears
	newSlice = newSlice[:2]
	printSlice(newSlice)

	// To remove all elements, simply set the slice to nil.
	// This will release the underlying array to the garbage collector
	// (assuming there are no other references). This is also the recommended way in Go Lang
	newSlice = nil
	printSlice(newSlice)
}

/*
----------------
[1 2 3 4 5]
[]int
5 5
----------------
----------------
[]
[]int
0 5
----------------
----------------
[1 2]
[]int
2 5
----------------
----------------
[]
[]int
0 0
----------------
*/
