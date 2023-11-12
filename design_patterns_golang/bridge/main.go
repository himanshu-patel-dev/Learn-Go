package main

import "fmt"

/*

------------- DEFINITION: BRIDGE PATTERN

Bridge is a structural design pattern that divides business logic
or huge class into separate class hierarchies that can be developed
independently.

------------- PROBLEM
Two shapes : Circle and Square
Two colors: Red and Blue

If we make each class for one combination (blue circle, red circle, blue sq, red sq)
then these are 4 combination/classes.

Rather we can use composition over inheritance
Circle (refer)-> blue/red class
Square (refer)-> blue/red class	

------------- APPLICATION
- Use the Bridge pattern when you want to divide and organize a monolithic class that 
has several variants of some functionality (for example, if the class can work with 
various database servers).

- Use the pattern when you need to extend a class in several orthogonal 
(independent) dimensions.

- Use the Bridge if you need to be able to switch implementations at runtime.

------------- IMPLEMENTATION

Abstraction (also called interface) is a high-level control layer for some entity. 
This layer isn’t supposed to do any real work on its own. It should delegate the 
work to the implementation layer (also called platform).

Abstraction: the GUI layer of the app.
Implementation: the operating systems’ APIs.

- In below example mac/pc class act as abstraction and printer class act as implementation

- Identify the orthogonal dimensions in your classes. These independent concepts 
could be: abstraction/platform, domain/infrastructure, 
front-end/back-end, or interface/implementation.

- See what operations the client needs and define them in the base abstraction class.

- Determine the operations available on all platforms. Declare the ones that the 
abstraction needs in the general implementation interface.

- For all platforms in your domain create concrete implementation classes, 
but make sure they all follow the implementation interface.

- Inside the abstraction class, add a reference field for the implementation type. 
The abstraction delegates most of the work to the implementation object that’s 
referenced in that field.

- If you have several variants of high-level logic, create refined abstractions 
for each variant by extending the base abstraction class.

- The client code should pass an implementation object to the abstraction’s constructor 
to associate one with the other. After that, the client can forget about the 
implementation and work only with the abstraction object.

*/

func main() {
	// here instead of 4 class 2 (computer) x 2 (printer)
	// we just managed to do this with 2 + 2 classes
	// this becomes more resonating when the n * m is too
	// large then n + m
	hpPrinter := &HP{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}

// Print request for mac
// Printing by a HP Printer

// Print request for mac
// Printing by a EPSON Printer

// Print request for windows
// Printing by a HP Printer

// Print request for windows
// Printing by a EPSON Printer
