package main

import "fmt"

// defining a interface
type Number interface {
	int64 | float64
}

func main() {
	// initialize a map for the integer values
	ints := map[string]int64{
		"first":  20,
		"second": 25,
	}
	// init a map for float values
	floats := map[string]float64{
		"first":  20.5,
		"second": 25.6,
	}
	// need to call different func based on types
	fmt.Println(sumInts(ints))     // 45
	fmt.Println(sumFloats(floats)) // 46.1

	// Generic Sums, type parameters inferred
	fmt.Println(sumIntsOrFloats(ints))   // 45
	fmt.Println(sumIntsOrFloats(floats)) // 46.1
	// can also be called as below in some places - Generic Sums
	// fmt.Println(sumIntsOrFloats[string, int64](ints))     // 45
	// fmt.Println(sumIntsOrFloats[string, float64](floats)) // 46.1

	// Generic Sums with Constraint:
	fmt.Println(sumNumbers(ints))
	fmt.Println(sumNumbers(floats))
}

// works only for int64
func sumInts(m map[string]int64) int64 {
	var sum int64
	for _, v := range m {
		sum += v
	}
	return sum
}

// works only for float64
func sumFloats(m map[string]float64) float64 {
	var sum float64
	for _, v := range m {
		sum += v
	}
	return sum
}

// works for both int and float
func sumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

// have type constraint along with generics
// Number interface takes out type union out of func declaration
func sumNumbers[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}
