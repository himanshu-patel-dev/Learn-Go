package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println() // i not accessible here

	i := 1
	for ; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println(i)
}
