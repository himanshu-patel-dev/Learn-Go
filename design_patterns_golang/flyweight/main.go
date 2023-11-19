package main

import "fmt"

/*

------------- DEFINITION: FLYWEIGHT
Flyweight is a structural design pattern that lets you fit more 
objects into the available amount of RAM by sharing common parts 
of state between multiple objects instead of keeping all of the 
data in each object.

------------- APPLICATION
Use the Flyweight pattern only when your program must support a 
huge number of objects which barely fit into available RAM.

------------- IMPLEMENTATION
Divide fields of a class that will become a flyweight into two parts:
- the intrinsic state: the fields that contain unchanging data 
  duplicated across many objects
- the extrinsic state: the fields that contain contextual data unique
  to each object
- Leave the fields that represent the intrinsic state in the class, 
  but make sure they’re immutable. They should take their initial 
  values only inside the constructor.

- Go over methods that use fields of the extrinsic state. For each 
  field used in the method, introduce a new parameter and use it 
  instead of the field.

- Optionally, create a factory class to manage the pool of flyweights. 
  It should check for an existing flyweight before creating a new one. 
  Once the factory is in place, clients must only request flyweights 
  through it. They should describe the desired flyweight by passing 
  its intrinsic state to the factory.

- The client must store or calculate values of the extrinsic state 
  (context) to be able to call methods of flyweight objects. For 
  the sake of convenience, the extrinsic state along with the 
  flyweight-referencing field may be moved to a separate context class.

*/

func main() {
	game := &Game{}

	//Add Terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	//Add CounterTerrorist
	game.addPolice(PoliceDressType)
	game.addPolice(PoliceDressType)
	game.addPolice(PoliceDressType)

	// even after making multiple instaces
	// of terrorist and police still we have only
	// 2 instaces of dresses becuase that 2 instances
	// are being referred among all the terrorist and police
	// instances
	for dressType, dress := range dressFactorySingleInstance.dressMap {
		fmt.Printf("DressColorType: %s\tDressColor: %s\n", dressType, dress.getColor())
	}
}

// DressColorType: TDress  DressColor: red
// DressColorType: PDress  DressColor: green
