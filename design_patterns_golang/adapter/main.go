package main

/*

use command : `go run .`
not `go run main.go`

---------------- DEFINITION: Adapter Pattern
Adapter is a structural design pattern,
which allows incompatible objects to collaborate.

---------------- APPLICABILITY
- Use the Adapter class when you want to use some existing class, 
but its interface isn’t compatible with the rest of your code.

- Use the pattern when you want to reuse several existing subclasses 
that lack some common functionality that can’t be added to the superclass.

*/

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(
		windowsMachineAdapter,
	)
}
