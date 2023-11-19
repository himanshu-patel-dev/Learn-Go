package main

import "fmt"

/*

use command : `go run .`
not `go run main.go`

--------------------- DEFINITION: DECORATOR PATTERN

Decorator is a structural pattern that allows adding new behaviors
to objects dynamically by placing them inside special wrapper objects,
called decorators.

Thus the calling method on decorators is exactaly same
as the method we need to call on base object

--------------------- DEFINITION: APPLICATIONS

Use the Decorator pattern when you need to be able to assign extra behaviors 
to objects at runtime without breaking the code that uses these objects.

Use the pattern when it’s awkward or not possible to extend an object’s 
behavior using inheritance.

*/

func main() {
	pizza := &VeggieMania{}
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}
	fmt.Printf(
		"Price of veggeMania with tomato and cheese topping is %d\n",
		pizzaWithCheeseAndTomato.getPrice(),
	)
	// Price of veggeMania with tomato and cheese topping is 32
}
