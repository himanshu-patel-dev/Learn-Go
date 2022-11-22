package main

import "fmt"

// normal func
func plus(a int, b int) int {
	return a + b
}

// args all of same type can have type defined in the end
func plusPlus(a, b, c int) int {
	return a + b + c
}

// multiple return value
func plusMult(a, b int) (int, int) {
	return a + b, a * b
}

// variadic function - any number of arguments
// Within the function, the type of nums is equivalent to []int.
func sum(nums ...int) {
	fmt.Printf("Type: %T \n", nums) // Type: []int
	total := 0

	for num := range nums {
		total += num
	}
	fmt.Println(total) // 6
}

// closures - anonymous function
func intSequence() func() int {
	// init i with some value which serves
	// as base for returned func
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println(plus(1, 2))        // 3
	fmt.Println(plusPlus(1, 2, 3)) // 6
	a, b := plusMult(1, 2)
	fmt.Println(a, b) // 3 2
	sum(1, 2, 3, 4)   // 6

	first_closure := intSequence()
	fmt.Println(first_closure()) // 1
	fmt.Println(first_closure()) // 2
	fmt.Println(first_closure()) // 3

	second_closure := intSequence()
	fmt.Println(second_closure()) // 1
}
