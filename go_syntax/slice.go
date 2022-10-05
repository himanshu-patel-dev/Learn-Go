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
	// slice literals are defined just like arrays
	// but without the element count
	slice := []int{1, 2, 3}
	printSlice(slice)

	integer := make([]int, 5, 10)
	printSlice(integer)

	for _, v := range integer {
		fmt.Print(v)
	}
	fmt.Println()

	// append to slice, len will inc but cap will remain same until
	// len become more than capacity
	integer = append(integer, 1, 2, 3, 4, 5)
	printSlice(integer) // 10 10

	// if len further inc then capacity will become double of prev
	integer = append(integer, 11)
	printSlice(integer) // 11 20
}
