package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Numeric Conversion
	i, _ := strconv.Atoi("-42")   // only for int, not for float
	fmt.Printf("%T : %d\n", i, i) // int : -42

	s := strconv.Itoa(-43)        // only for int, not for float
	fmt.Printf("%T : %s\n", s, s) // string : -43

	b, _ := strconv.ParseBool("true")
	fmt.Printf("%T : %t\n", b, b) // bool : true

	f, _ := strconv.ParseFloat("3.14", 64) // float string, bit size
	fmt.Printf("%T : %f\n", f, f)          // float : 3.14

	// parse to decimal
	j, _ := strconv.ParseInt("32", 10, 64) // int string, base, bit size
	fmt.Printf("%T : %d\n", j, j)          // int64 : 32

	// parse to binary
	j, _ = strconv.ParseInt("101", 2, 64) // int string, base, bit size
	fmt.Printf("%T : %d\n", j, j)         // int64 : 5
}
