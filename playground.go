package main

import "fmt"

func main() {
	var intVar int = 5
	var floatVar float32 = 45.5
	var res1 = float32(intVar) + floatVar
	var res2 = intVar + int(floatVar)
	fmt.Println(res1, res2)
}
