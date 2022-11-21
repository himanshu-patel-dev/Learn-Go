package main

import (
	"fmt"
)

func main() {
	anArray := [5]int{0, 1, -1, 2, -2}
	fmt.Println(anArray)      // [0 1 -1 2 -2]
	fmt.Println(len(anArray)) // 5

	twoD := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(twoD) // [[1 2 3 4] [5 6 7 8] [9 10 11 12]]

	threeD := [3][2][1]int{ // [[[1] [2]] [[3] [4]] [[5] [6]]]
		{
			{1}, {2},
		}, {
			{3}, {4},
		}, {
			{5}, {6},
		},
	}
	fmt.Println(threeD)

	// printing 3D array
	fmt.Println("----------------")
	for i := 0; i < len(threeD); i++ {
		for j := 0; j < len(threeD[i]); j++ {
			for k := 0; k < len(threeD[i][j]); k++ {
				fmt.Print(threeD[i][j][k], " ")
			}
		}
	}
	// 1 2 3 4 5 6

	fmt.Println("\n----------------")
	for _, v := range threeD {
		for _, m := range v {
			for _, s := range m {
				fmt.Print(s, " ")
			}
		}
	}
	// 1 2 3 4 5 6
	fmt.Println("\n----------------")
}
