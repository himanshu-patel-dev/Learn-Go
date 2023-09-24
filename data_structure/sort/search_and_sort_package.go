package main

import (
	"fmt"
	"math"
	"sort"
)

func printSearchResultInts(i int, target int, list []int) {
	fmt.Printf("Index for insertion/first occurace of target: %d\n", i)
	if i < len(list) {
		fmt.Printf("Target Found: %v\n", list[i] == target)
	} else {
		fmt.Printf("Target Found: %v\n", false)
	}
	fmt.Println("--------------------------------------")
}

func printSearchResultFloats(i int, target float64, list []float64) {
	fmt.Printf("Index for insertion/first occurace of target: %d\n", i)
	if i < len(list) {
		fmt.Printf("Target Found: %v\n", list[i] == target)
	} else {
		fmt.Printf("Target Found: %v\n", false)
	}
	fmt.Println("--------------------------------------")
}

func printSearchResultString(i int, target string, list []string) {
	fmt.Printf("Index for insertion/first occurace of target: %d\n", i)
	if i < len(list) {
		fmt.Printf("Target Found: %v\n", list[i] == target)
	} else {
		fmt.Printf("Target Found: %v\n", false)
	}
	fmt.Println("--------------------------------------")
}

func BinarySearch() {
	numSlice := []int{1, 3, 5, 7, 9, 9, 11, 13}
	target, n := 8, len(numSlice)

	i := sort.Search(n, func(i int) bool {
		return numSlice[i] >= target
	})
	// first index of 8, first bigger element than target
	// binary search = O(log(n))
	printSearchResultInts(i, target, numSlice)

	target, n = 15, len(numSlice)
	i = sort.Search(n, func(i int) bool {
		return numSlice[i] >= target
	})
	printSearchResultInts(i, target, numSlice)

	target, n = 9, len(numSlice)
	i = sort.Search(n, func(i int) bool {
		return numSlice[i] >= target
	})
	printSearchResultInts(i, target, numSlice)

	// Index for insertion/first occurace of target: 4
	// Target Found: false
	// --------------------------------------
	// Index for insertion/first occurace of target: 8
	// Target Found: false
	// --------------------------------------
	// Index for insertion/first occurace of target: 4
	// Target Found: true
	// --------------------------------------
}

func BinarySearchFloats() {
	fmt.Println("Search Floats")

	list := []float64{1.0, 2.0, 3.3, 4.6, 6.1, 7.2, 8.0}
	target := 2.0
	i := sort.SearchFloat64s(list, target)
	printSearchResultFloats(i, target, list)

	target = 0.5
	i = sort.SearchFloat64s(list, target)
	printSearchResultFloats(i, target, list)

	// Index for insertion/first occurace of target: 1
	// Target Found: true
	// --------------------------------------
	// Index for insertion/first occurace of target: 4
	// Target Found: false
}

func BinarySearchInts() {
	fmt.Println("Search Ints")

	list := []int{1, 2, 3, 4, 6, 7, 8}

	target := 2
	i := sort.SearchInts(list, target)
	printSearchResultInts(i, target, list)

	target = 5
	i = sort.SearchInts(list, target)
	printSearchResultInts(i, target, list)

	// Index for insertion/first occurace of target: 1
	// Target Found: true
	// --------------------------------------
	// Index for insertion/first occurace of target: 0
	// Target Found: false
}

func BinarySearchString() {
	fmt.Println("Search Strings")
	list := []string{"a", "c", "e", "g", "i"}

	target := "e"
	i := sort.SearchStrings(list, target)
	printSearchResultString(i, target, list)

	target = "b"
	i = sort.SearchStrings(list, target)
	printSearchResultString(i, target, list)

	// Search Strings
	// Index for insertion/first occurace of target: 2
	// Target Found: true
	// --------------------------------------
	// Index for insertion/first occurace of target: 1
	// Target Found: false
}

func SortFloats() {
	// inplace sort, copy slice in order to present original one
	floatSlice := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	sort.Float64s(floatSlice)
	fmt.Println(floatSlice)
	// [-3.8 -1.3 0.7 2.6 5.2]

	floatSlice = []float64{math.Inf(1), math.NaN(), math.Inf(-1), 0.0} // unsorted
	sort.Float64s(floatSlice)
	fmt.Println(floatSlice)
	// [NaN -Inf 0 +Inf]

	fmt.Println(sort.Float64sAreSorted(floatSlice)) // true
	fmt.Println("----------------------------------------------")
}

func SortInts() {
	// inplace sort, make a copy to retain original
	intSlice := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(intSlice)
	fmt.Println(intSlice) // [5, 2, 6, 3, 1, 4]

	fmt.Println(sort.IntsAreSorted(intSlice)) // true
	fmt.Println("----------------------------------------------")
}

type Person struct {
	Name string
	Age  int
}

func SortStruct() {
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
	cmp := func(i, j int) bool {
		return people[i].Age < people[j].Age
	}
	sort.Slice(people, cmp)
	fmt.Println("After Sorting: ", people)
	fmt.Println("After Sorting: ", people_copy)
	fmt.Println("Sorted Slice: ", sort.SliceIsSorted(people, cmp))
	// Before Sorting:  [{Bob 31} {John 42} {Michael 17} {Jenny 26}]
	// After Sorting:  [{Michael 17} {Jenny 26} {Bob 31} {John 42}]
	// After Sorting:  [{Bob 31} {John 42} {Michael 17} {Jenny 26}]
	fmt.Println("----------------------------------------------")

}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (by ByAge) Len() int           { return len(by) }
func (by ByAge) Swap(i, j int)      { by[i], by[j] = by[j], by[i] }
func (by ByAge) Less(i, j int) bool { return by[i].Age < by[j].Age }

func SortInterface() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	interfaceSlice := ByAge(people)
	sort.Sort(interfaceSlice)
	fmt.Println(interfaceSlice)
	// [{Michael 17} {Jenny 26} {Bob 31} {John 42}]
	// to use IsSorted the slice should have impremented the Interface for sort
	fmt.Println(sort.IsSorted(interfaceSlice)) // true
	fmt.Println("----------------------------------------------")
}

func SortStrings() {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)                        // [Alpha Bravo Delta Go Gopher Grin]
	fmt.Println(sort.StringsAreSorted(s)) // true
	fmt.Println("----------------------------------------------")
}

func main() {
	// search
	BinarySearch()
	BinarySearchInts()
	BinarySearchFloats()
	BinarySearchString()

	// sort
	SortFloats()
	SortInts()
	SortStruct()
	SortInterface()
	SortStrings()
}
