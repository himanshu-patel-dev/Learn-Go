package main

import "fmt"

func main() {
	// infinite loop (if no break condition)

	// for {
	// 	fmt.Println(1)
	// }

	// while loop
	i := 1
	for {
		fmt.Print(i, " ")
		if i == 5 {
			break
		}
		i++
	}
}
