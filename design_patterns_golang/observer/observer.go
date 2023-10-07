package main

import "fmt"

// Observer is the object which is observing a
// Subject and get notified on it's update
type Observer interface {
	update(string)
	getID() string
}

// Customer is implementation of Observer
type Customer struct {
	id string
}

func newCustomer(email string) *Customer {
	return &Customer{id: email}
}

func (c *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *Customer) getID() string {
	return c.id
}
