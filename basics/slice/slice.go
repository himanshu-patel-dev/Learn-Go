package main

import "fmt"

// ### The shortcomings of Go arrays
// - Once you define an array, you cannot change its size, which means that Go arrays are not dynamic.
// Putting it simply, if you need to add an element to an existing array that has no space left,
// you will need to create a bigger array and copy all of the elements of the old array to the new one.

// - Second, when you pass an array to a function as a parameter, you actually pass a copy of the array,
// which means that any changes you make to an array inside a function will be lost after the function exits.

// - Last, passing a large array to a function can be pretty slow, mostly because Go has to create a
// copy of the array. The solution to all of these problems is to use Go `slices`.

// Additionally, passing a big slice to a function is significantly faster than passing
// an array with the same number of elements because Go will not have to make a copy of the
// slice—it will just pass the memory address of the slice variable.

func main() {
	// slice literals are defined just like arrays
	// but without the element count
	slice := []int{1, 2, 3}
	fmt.Println(slice) // [1 2 3]

	// Pointer: The pointer is used to points to the first element of the array that
	// is accessible through the slice. Here, it is not necessary that the pointed
	// element is the first element of the array.

	// Length: The length is the total number of elements present in the array.

	// Capacity: The capacity represents the maximum size upto which it can expand.

	integer := make([]int, 5, 10)           // slice type, size, capacity
	fmt.Println(integer)                    // [0 0 0 0 0] initialized with zero value of type
	fmt.Println(len(integer), cap(integer)) // 5 10

	// append to slice, len will inc but cap will remain same until
	// len become more than capacity
	integer = append(integer, 1, 2, 3, 4, 5)
	fmt.Println(integer)                    // [0 0 0 0 0 1 2 3 4 5]
	fmt.Println(len(integer), cap(integer)) // 10 10

	// if len further inc then capacity will become double of prev
	integer = append(integer, 11)
	fmt.Println(integer)                    // [0 0 0 0 0 1 2 3 4 5 11]
	fmt.Println(len(integer), cap(integer)) // 11 20

	// slice shallow copy
	s1 := make([]int, 5)
	reSlice := s1[1:3]
	fmt.Println(s1)      // [0 0 0 0 0]
	fmt.Println(reSlice) // [0 0]
	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println(s1)      // [0 -100 123456 0 0]
	fmt.Println(reSlice) // [-100 123456]

	// Type of slice
	fmt.Printf("Type: %T \n", reSlice) // Type: []int

	newSlice := []int{1, 2, 3, 4, 5}

	// keep allocated memory - slice the array to zero length
	// but the capacity and type remians the same
	newSlice = newSlice[:0]
	fmt.Println("Slice: ", newSlice)
	// Slice: []

	// if slice is extended the original data reappears
	// Because both slices point to the same memory address.
	// re-slice process does not make a copy of the original slice.
	newSlice = newSlice[:2]
	fmt.Println("Slice: ", newSlice)
	// Slice: [1 2]

	// To remove all elements, simply set the slice to nil.
	// This will release the underlying array to the garbage collector
	// (assuming there are no other references). This is also the recommended way in Go Lang
	newSlice = nil
	fmt.Println("Slice: ", newSlice)
	// Slice: []

	// The second problem of re-slicing is that, even if you re-slice a slice in order
	// to use a small part of the original slice, the underlying array from the original
	// slice will be kept in memory for as long as the smaller re-slice exists because
	// the original slice is being referenced by the smaller re-slice. Although this
	// is not truly important for small slices, it can cause problems when you are
	// reading big files into slices and you only want to use a small part of them.

	// use `[:]` notation for creating a new slice from an existing slice or array

	// make a slice with just outer slice len
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		// dynamically assign the sub slice
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // 2d:  [[0] [1 2] [2 3 4]]
}
