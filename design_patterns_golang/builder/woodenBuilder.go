package main

type WoodenBuilder struct {
	House
}

func (w *WoodenBuilder) setWindowType() {
	w.windowType = "wooden window"
}

func (w *WoodenBuilder) setDoorType() {
	w.doorType = "wooden door"
}

func (w *WoodenBuilder) setNumFloor() {
	w.floor = 2
}

func (w *WoodenBuilder) getHouse() House {
	return House{
		doorType:   w.doorType,
		windowType: w.windowType,
		floor:      w.floor,
	}
}

func newWoodenBuilder() *WoodenBuilder {
	return &WoodenBuilder{}
}
