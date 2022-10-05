package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		}
		if i == 5 {
			fmt.Println("Breaking out of loop")
			break // break here
		}
		fmt.Println("The value of i is", i)
	}
	fmt.Println("Exiting program")
}

/*
The value of i is 0
The value of i is 1
The value of i is 3
The value of i is 4
Breaking out of loop
Exiting program
*/
