package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name   string
	height int
	weight int
}

func main() {
	// sorting string array
	strs := []string{"c", "b", "a"}
	sort.Strings(strs)
	fmt.Println("String: ", strs) // String:  [a b c]

	// sorting int array
	ints := []int{4, 2, 1}
	sort.Ints(ints)
	fmt.Println(ints) // [1 2 4]

	// sorting custom array
	mySlice := make([]Person, 0)
	mySlice = append(mySlice, Person{"Mihalis", 180, 90})
	mySlice = append(mySlice, Person{"Bill", 134, 45})
	mySlice = append(mySlice, Person{"Marietta", 155, 45})
	mySlice = append(mySlice, Person{"Epifanios", 144, 50})
	mySlice = append(mySlice, Person{"Athina", 134, 40})
	fmt.Println("Slice : ", mySlice)
	// Slice :  [{Mihalis 180 90} {Bill 134 45} {Marietta 155 45} {Epifanios 144 50} {Athina 134 40}]

	// lesser height first
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)
	// <: [{Bill 134 45} {Athina 134 40} {Epifanios 144 50} {Marietta 155 45} {Mihalis 180 90}]

	// more height first
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)
	// >: [{Mihalis 180 90} {Marietta 155 45} {Epifanios 144 50} {Bill 134 45} {Athina 134 40}]
}
