package main

import "fmt"

func main() {
	// mind the , on the last line before the closing }
	mapping := map[int]string{
		1: "First",
		2: "Second",
		3: "Third",
	}
	fmt.Println(mapping)

	// create map without initializing
	maps := make(map[int]string)
	// now assign key value paris
	maps[1] = "First"
	fmt.Println(maps)

	// delete a key from map
	delete(mapping, 3)

	// iterate over map
	for key, value := range mapping {
		fmt.Println(key, value)
	}

	// checking if a element exists in a map
	// event if the key don't exists we get the zero value of
	// data type define for values in map
	val, ok := mapping[2]
	fmt.Println("Val and ok: ", val, ok)

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

	// Map are pass by reference in function call, i.e. if map is passed
	// to a function which end up modifing the map then the actual map is
	// modified not a copy value
	m := make(map[string]int)
	m["hello"] = 1
	m["world"] = 2
	fmt.Println(m) // map[hello:1 world:2]
	helloWorld(m)
	fmt.Println(m) // map[hello:100 world:2]
}

func helloWorld(m map[string]int) {
	m["hello"] = 100
}

/*
map[1:First 2:Second 3:Third]
map[1:First]
1 First
2 Second
Val and ok:  Second true
map[]
Lookup:  0 false
map[]
nil Map
*/
