package main

import "fmt"

/*

use command : `go run .`
not `go run main.go`

DEFINITION: FACTORY PATTERN
lets you produce families of related objects without specifying their concrete classes.

*/

// choose factory based on string input - to decide between all the factories
// available we have to make decision based on something
func GetSportsFactory(brand string) (InterfaceSportFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}

func main() {
	// get the factory based on name
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	// once we get the factory then getting the
	// product of that particular factory should be
	// just a method call
	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

// Logo: nike
// Size: 15
// Logo: nike
// Size: 15
// Logo: adidas
// Size: 14
// Logo: adidas
// Size: 14
