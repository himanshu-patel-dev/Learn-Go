package main

import "fmt"

// You should be very careful when using the `copy()` function
// on slices because the built-in `copy(dst, src)` copies the
// minimum of the `len(dst)` and `len(src)` elements.

// As `copy()` only accepts slice arguments, you should also use
// the `[:]` notation to convert the array into a slice. If you
// try to copy an array into a slice or vice versa without using
// the `[:]` notation, the program will fail to compile and will
// display one of the following error messages

// ./a.go:42:6: first argument to copy should be slice; have [5]int
// ./a.go:43:6: second argument to copy should be slice or string; have
// [5]int
// ./a.go:44:6: arguments to copy must be slices; have [5]int, [5]int

func main() {

	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6) // [-10 1 2 3 4 5]
	fmt.Println("a4:", a4) // [-1 -2 -3 -4]

	copy(a6, a4)           // dest src
	fmt.Println("a6:", a6) // [-1 -2 -3 -4 4 5]
	fmt.Println("a4:", a4) // [-1 -2 -3 -4]
	fmt.Println()

	b6 := []int{-10, 1, 2, 3, 4, 5}
	b4 := []int{-1, -2, -3, -4}
	copy(b4, b6)
	fmt.Println("b6:", b6) // [-10 1 2 3 4 5]
	fmt.Println("b4:", b4) // [-10 1 2 3 ]

	fmt.Println()
	array4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, 1, -1, -1, 5, -5}
	copy(s6, array4[1:])
	fmt.Println("array4:", array4[:]) // [4 -4 4 -4]
	fmt.Println("s6:", s6)            // [-4 4 -4 -1 5 -5]

	fmt.Println()

	array5 := [5]int{5, -5, 5, -5, 5}
	s7 := []int{7, 7, -7, -7, 7, -7, 7}
	copy(array5[1:], s7)
	fmt.Println("array5:", array5) // [5 7 7 -7 -7]
	fmt.Println("s7:", s7)         // [7 7 -7 -7 7 -7 7]

}
