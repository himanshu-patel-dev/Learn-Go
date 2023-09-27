package main

type Director struct {
	builder IBuilder
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}
