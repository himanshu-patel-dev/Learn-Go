package main

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "wooden" {
		return newWoodenBuilder()
	}

	if builderType == "metallic" {
		return newMetallicBuilder()
	}
	return nil
}
