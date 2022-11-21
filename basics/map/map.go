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
	// first 1
	// second 2

	// checking if a element exists in a map
	// event if the key don't exists we get the zero value of
	// data type define for values in map
	ele, ok := firstMap["third"]
	fmt.Println("Element exists: ", ele, ok) // 0 false

	// making a nil map
	bMap := map[string]int{}
	bMap = nil
	fmt.Println(bMap) // map[]

	// lookup do not give any error, similarly range and len
	ele2, ok2 := bMap["key"]
	fmt.Println("Lookup: ", ele2, ok2) // Lookup:  0 false

	// inserting in nil map gives error
	// bMap["test"] = 1

	// panic: assignment to entry in nil map
	// goroutine 1 [running]:
	// main.main()
	// 	~/Learn-Go/basics/map/map.go:49 +0x505
	// exit status 2

	// how a nil map get into picture
	// when we declare a map but forget to initialize it with make()
	var cMap map[string]int
	fmt.Println(cMap) // map[]
	if cMap == nil {  // nil Map
		fmt.Println("nil Map")
	} else {
		fmt.Println("Not nil Map")
	}
	// nil Map

}
