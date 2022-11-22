package main

import "fmt"

type person struct {
	name string
	age  int
}

// if we send
func newPerson(name string, age int) *person {
	p := person{name: name}
	p.age = age
	return &p
}

func main() {
	// way to instantiate
	fmt.Println(person{"Bob", 20})              // {Bob 20}
	fmt.Println(person{name: "Alice", age: 30}) // {Alice 30}
	// zero value
	fmt.Println(person{name: "Fred"}) // {Fred 0}
	fmt.Println(person{age: 10})      // { 10}
	// pointer
	fmt.Println(&person{name: "Ann", age: 40}) // &{Ann 40}
	fmt.Println(newPerson("Jon", 50))          // &{Jon 50}
}
