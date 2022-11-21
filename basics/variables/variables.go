package main

import "fmt"

func main() {
	// declare one or more variable using var
	var a = "first"
	fmt.Println(a) // first

	// declare multiple var of same type and initialize them
	// var b, c = 1, 2 // this also works
	// all var in should have same type
	var b, c int = 1, 2
	fmt.Println(b, c) // 1 2

	d := "second"
	e := true
	f := 5
	fmt.Println(d, e, f) // second true 5
}
