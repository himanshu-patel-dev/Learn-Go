package main

import "fmt"

// Subject interface need to be implemented by every
// item acting as a point of observation
type Subject interface {
	register(Observer)
	deregister(Observer)
	notifyAll()
}

// Item is a implementation of Subject
type Item struct {
	name         string
	inStock      bool
	observerList []Observer
}

// newItem is constructor for Item
func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) deregister(o Observer) {
	i.observerList = removeFromList(i.observerList, o)
}

func removeFromList(observerList []Observer, observerToRemove Observer) []Observer {
	n := len(observerList)
	for i, observer := range observerList {
		if observer.getID() == observerToRemove.getID() {
			observerList[i], observerList[n-1] = observerList[n-1], observerList[i]
			observerList = observerList[:n-1]
			return observerList
		}
	}
	// if the object to remove is not present
	// return the list as is
	return observerList
}

func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}
