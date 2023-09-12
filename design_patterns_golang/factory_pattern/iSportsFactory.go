package main

type InterfaceSportFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}
