package main

import "fmt"

/*

use command : `go run .`
not `go run main.go`

DEFINITION: Builder Pattern
The Builder pattern is used when the desired product
is complex and requires multiple steps to complete.

*/

func main() {
	woodenBuilder := newWoodenBuilder()
	metallicBuilder := newMetallicBuilder()

	// director direct wooden builder for house
	director := newDirector(woodenBuilder)
	woodenHouse := director.buildHouse()
	printHouse((woodenHouse))

	// director direct metallic builder for house
	director.setBuilder(metallicBuilder)
	metallicHouse := director.buildHouse()
	printHouse(metallicHouse)
}

func printHouse(house House) {
	fmt.Printf("Normal House Door Type: %s\n", house.doorType)
	fmt.Printf("Normal House Window Type: %s\n", house.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", house.floor)
	fmt.Println("-----------------------------")
}

// Normal House Door Type: wooden door
// Normal House Window Type: wooden window
// Normal House Num Floor: 2
// -----------------------------
// Normal House Door Type: metallic door
// Normal House Window Type: metallic window
// Normal House Num Floor: 3
// -----------------------------
