package main

import "fmt"

func main() {
	// works fine
	aMap := map[string]int{}
	aMap["test"] = 1

	// making a nil map
	bMap := map[string]int{}
	bMap = nil
	fmt.Println(bMap)

	// lookup do not give any error, similarly range and len
	ele, ok := bMap["key"]
	fmt.Println("Lookup: ", ele, ok)

	// inserting in nil map gives error
	// bMap["test"] = 1

	// how a nil map get into picture
	var cMap map[string]int
	fmt.Println(cMap) // map[]
	if cMap == nil {  // nil Map
		fmt.Println("nil Map")
	} else {
		fmt.Println("Not nil Map")
	}
}
