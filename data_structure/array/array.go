package main

import "fmt"

func main() {
	// arrays
	arr := [5]int{2, 4, 6, 8, 10}
	fmt.Println(arr)

	for i, val := range arr {
		fmt.Println("index: ", i, " value: ", val)
	}
	fmt.Println("\n----------------")

	arr = [5]int{1, 3, 4, 5, 2} // can reassign with same size different array
	fmt.Println("Reassigned Array: ", arr)
	fmt.Println("\n----------------")

	// arr = [6]int{1, 3, 4, 5, 2, 0} // cannot reassign with different size array

	// slice
	slice := []int{1, 2, 3, 4}
	fmt.Println("Slice: ", slice, " Capacity of slice: ", cap(slice), " Len of slice: ", len(slice))
	slice = append(slice, 8) // capacity doubles to accomodate future elements possibility
	fmt.Println("Slice: ", slice, " Capacity of slice: ", cap(slice), " Len of slice: ", len(slice))
	fmt.Println("\n----------------")

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
	fmt.Println("\n----------------")
	for _, v := range threeD {
		for _, m := range v {
			for _, s := range m {
				fmt.Print(s, " ")
			}
		}
	}
	fmt.Println("\n----------------")

}
