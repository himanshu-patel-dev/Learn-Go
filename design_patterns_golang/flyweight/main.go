package main

import "fmt"

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
