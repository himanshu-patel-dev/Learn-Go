package main

import (
	"fmt"
)

func main() {
	vendingMachine := newVendingMachine(1, 10)

	// select item
	err := vendingMachine.requestItem()
	checkError(err)
	// put money
	err = vendingMachine.insertMoney(10)
	checkError(err)
	// take the item
	err = vendingMachine.dispenseItem()
	checkError(err)

	fmt.Printf("\nCurrent State: %v\n\n", vendingMachine.getState())

	// no more item left, trying to take one
	// result in error
	err = vendingMachine.requestItem()
	checkError(err)
	err = vendingMachine.dispenseItem()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

//Item requestd
//Money entered is ok
//Dispensing Item
//
//Current State: NoItemState
//
//item out of stock
//item out of stock
