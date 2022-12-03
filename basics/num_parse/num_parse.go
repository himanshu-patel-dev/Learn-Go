package main

import (
	"fmt"
	"strconv"
)

// he built-in package strconv provides the number parsing.

func main() {
	// With ParseFloat, this 64 tells how many bits of precision to parse.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f) // 1.234 -- float64

	// For ParseInt, the 0 means infer the base from the string.
	// 64 requires that the result fit in 64 bits.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i) // 123 -- int64

	// ParseInt will recognize hex-formatted numbers
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d) // 456 -- int64

	// A ParseUint is also available
	u, _ := strconv.ParseInt("789", 0, 64)
	fmt.Println(u) // 789 -- int64

	// Atoi is a convenience function for basic base-10 int parsing.
	k, _ := strconv.Atoi("135")
	fmt.Println(k) // type of k is int and not string anymore

	// Parse functions return an error on bad input.
	_, e := strconv.Atoi("IamNumber")
	fmt.Println(e)
	// strconv.Atoi: parsing "IamNumber": invalid syntax
}
