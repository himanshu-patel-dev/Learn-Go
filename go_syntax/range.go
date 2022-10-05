package main

import "fmt"

func main() {
	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index:", i, "value: ", value)
	}
	fmt.Println("--------------")
	for _, value := range anArray {
		fmt.Println("value: ", value)
	}
	fmt.Println("--------------")
	for index := range anArray {
		fmt.Println("index: ", index)
	}
}
