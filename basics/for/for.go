package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	} // i not accessable here
	// 1
	// 2
	// 3

	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}
	// 1
	// 2
	// 3

	for { // an infinite loop
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	// 1 3 5
	// n not accessable here

	// infinite loop (if no break condition)

	// for {
	// 	fmt.Println(1)
	// }

	// while loop
	k := 1
	for {
		fmt.Print(k, " ")
		if k == 5 {
			break
		}
		k++
	}
	// 1 2 3 4 5
}
