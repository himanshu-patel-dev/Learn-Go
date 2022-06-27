package main

import "fmt"

func main() {
	// define map using make
	firstMap := make(map[string]int)
	firstMap["first"] = 1
	firstMap["second"] = 2
	fmt.Println("Map: ", firstMap) // Map:  map[first:1 second:2]

	// define map using map literal
	secondMap := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(secondMap) // map[a:1 b:2]

	// try deleting one element
	delete(secondMap, "a")
	fmt.Println(secondMap) // map[b:2]
	// try deleting again
	delete(secondMap, "a")
	fmt.Println(secondMap) // map[b:2]

	// iterate over all elements of map
	for key, value := range firstMap {
		fmt.Println(key, value)
	}

	// checking if a element exists in a map
	// event if the key don't exists we get the zero value of
	// data type define for values in map
	ele, ok := firstMap["third"]
	fmt.Println("Element exists: ", ele, ok) // 0 false
}
