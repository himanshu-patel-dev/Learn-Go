package main

import "fmt"

/*

use command : `go run .`
not `go run main.go`

---------------- DEFINITION: Builder Pattern
Builder pattern is used when the desired product
is complex and requires multiple steps to complete.
It let you create a complex object step by step.

---------------- APPLICATION
- Use the Builder pattern to get rid of a “telescoping constructor”.
    Pizza(int size) { ... }
    Pizza(int size, boolean cheese) { ... }
    Pizza(int size, boolean cheese, boolean pepperoni) { ... }

- Use the Builder pattern when you want your code to be able to create 
different representations of some product (stone and wooden houses).

---------------- IMPLEMENTATION
- clearly define the common construction steps for building all available 
product representations. Only those which are common among all builders.
Declare these steps in the base builder interface.

- Create a concrete builder class for each of the product representations 
and implement their construction steps.

Don’t forget about implementing a method for fetching the result of the construction. 
The reason why this method can’t be declared inside the builder interface is that 
various builders may construct products that don’t have a common interface. 
Therefore, you don’t know what would be the return type for such a method. 

- Think about creating a director class. It may encapsulate various ways to 
construct a product using the same builder object.

- The client code creates both the builder and the director objects. 
Before construction starts, the client must pass a builder object to the director. 
Usually, the client does this only once, via parameters of the director’s class 
constructor. The director uses the builder object in all further construction. 

- The construction result can be obtained directly from the director only if all 
products follow the same interface. Otherwise, the client should fetch the result 
from the builder.

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
