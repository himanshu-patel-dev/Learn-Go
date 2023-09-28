package main

/*

use command : `go run .`
not `go run main.go`

DEFINITION: Adapter Pattern
Adapter is a structural design pattern,
which allows incompatible objects to collaborate.

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
