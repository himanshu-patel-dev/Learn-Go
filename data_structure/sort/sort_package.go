package main

import (
	"fmt"
	"sort"
)

func SortStruct() {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println("Before Sorting: ", people)
	// closure function ask is element at i is
	// Less than element at j, if it return true
	// element at i will be before element at j in result
	// if it is false then i will be after j, kind of
	// reverse sort

	// inplace sort, array get modified, save a copy to see diff
	people_copy := make([]Person, len(people))
	copy(people_copy, people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("After Sorting: ", people)
	fmt.Println("After Sorting: ", people_copy)
	// Before Sorting:  [{Bob 31} {John 42} {Michael 17} {Jenny 26}]
	// After Sorting:  [{Michael 17} {Jenny 26} {Bob 31} {John 42}]
	// After Sorting:  [{Bob 31} {John 42} {Michael 17} {Jenny 26}]

}

func BinarySearch() {
	numSlice := []int{1, 3, 5, 7, 9, 9, 11, 13}
	target, n := 8, len(numSlice)
	printResult := func(i int, target int, list []int) {
		fmt.Printf("Index for insertion/first occurace of target: %d\n", i)
		if i < n {
			fmt.Printf("Target Found: %v\n", list[i] == target)
		} else {
			fmt.Printf("Target Found: %v\n", false)
		}
		fmt.Println("--------------------------------------")
	}

	i := sort.Search(n, func(i int) bool {
		return numSlice[i] >= target
	})
	// first index of 8, first bigger element than target
	// binary search = O(log(n))
	printResult(i, target, numSlice)

	target, n = 15, len(numSlice)
	i = sort.Search(n, func(i int) bool {
		return numSlice[i] >= target
	})
	printResult(i, target, numSlice)

	target, n = 9, len(numSlice)
	i = sort.Search(n, func(i int) bool {
		return numSlice[i] >= target
	})
	printResult(i, target, numSlice)

}

func main() {
	SortStruct()
	BinarySearch()
}
