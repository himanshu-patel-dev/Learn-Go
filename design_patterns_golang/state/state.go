package main

import "fmt"

/*
Consider each state as state of one finite
automata and all the methods on it are the branches
which takes that state to another states if possible

all methods are the possible action a user can take
while being in that state
*/

// interface

type State interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
	getState() string
}

// concrete implementation

// -------------------------------------------------------------

// NoItemState state
type NoItemState struct {
	vendingMachine *VendingMachine
}

func (i *NoItemState) requestItem() error {
	return fmt.Errorf("item out of stock")
}

func (i *NoItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *NoItemState) dispenseItem() error {
	return fmt.Errorf("item out of stock")
}

func (i *NoItemState) getState() string {
	return fmt.Sprintf("NoItemState")
}

// -------------------------------------------------------------

// HasItemState state
type HasItemState struct {
	vendingMachine *VendingMachine
}

func (i *HasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("no item present")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *HasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (i *HasItemState) dispenseItem() error {
	return fmt.Errorf("please select item first")
}

func (i *HasItemState) getState() string {
	return fmt.Sprintf("HasItemState")
}

// -------------------------------------------------------------

// ItemRequestedState state
type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("item dispence in progress")
}

func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	// return changes if extra
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("please insert money first")
}

func (i *ItemRequestedState) getState() string {
	return fmt.Sprintf("ItemRequestedState")
}

// -------------------------------------------------------------

// HasMoneyState state
type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) requestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (i *HasMoneyState) addItem(count int) error {
	return fmt.Errorf("item dispence in progress")
}

func (i *HasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("money alredy deposited")
}

func (i *HasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}

func (i *HasMoneyState) getState() string {
	return fmt.Sprintf("hasMoneyState")
}
