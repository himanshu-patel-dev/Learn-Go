package main

import "fmt"

/*

use command : `go run .`
not `go run main.go`

---------------- DEFINITION: FACTORY PATTERN
lets you produce objects of one particular family 
without specifying their concrete classes.

---------------- APPLICATION
- use factory method when your code don't know the exact type and
dependencies of objects your code should work with

- factory method separate the construction code from the code that
uses it. thus it's easier to update product construction code independently
from the client code

to add a new product type to the app, you’ll only need to create a new 
creator subclass and override the factory method in it.

- use it to provide user with a way to extend internal component of library/framework

Imagine that you write an app using an open source UI framework.
You extend the standard Button class with a glorious RoundButton subclass.
But now you need to tell the main UIFramework class to use the new button 
subclass instead of a default one.

you create a subclass UIWithRoundButtons from a base framework class and override 
its createButton method. While this method returns Button objects in the base 
class, you make your subclass return RoundButton objects. Now use the 
UIWithRoundButtons class instead of UIFramework. 

---------------- HOW TO IMPLEMENT

- Make all products follow the same interface. 
This interface should declare methods that make sense in every product.

- Make each factory follow the factory interface to make creator class 
consistent with code, it can create any object with same methods. 

- return factory classes from inside the creator class.
The return type of the factory should match the common factory interface.
Creator class is the one which create factory, here GetSportsFactory.

*/

// choose factory based on string input - to decide between all the factories
// available we have to make decision based on something
func GetSportsFactory(brand string) (InterfaceSportFactory, error) {
	if brand == "adidas" {
		return &AdidasFactory{}, nil
	}

	if brand == "nike" {
		return &NikeFactory{}, nil
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
