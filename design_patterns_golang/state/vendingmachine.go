package main

import "fmt"

type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State
	// this points to current state
	currentState State
	// this is the data which determine each state
	// along with the request performed
	itemCount int
	itemPrice int
}

func newVendingMachine(itemCount, itemPrice int) *VendingMachine {
	// make new machine with the initial data
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	// attach machine to states
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}
	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *VendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

func (v *VendingMachine) getState() string {
	return v.currentState.getState()
}

func (v *VendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}
