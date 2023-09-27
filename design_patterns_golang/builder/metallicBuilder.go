package main

type MetallicBuilder struct {
	House
}

func (m *MetallicBuilder) setDoorType() {
	m.doorType = "metallic door"
}

func (m *MetallicBuilder) setWindowType() {
	m.windowType = "metallic window"
}

func (m *MetallicBuilder) setNumFloor() {
	m.floor = 3
}

func (m *MetallicBuilder) getHouse() House {
	return House{
		doorType:   m.doorType,
		windowType: m.windowType,
		floor:      m.floor,
	}
}

func newMetallicBuilder() *MetallicBuilder {
	return &MetallicBuilder{}
}
