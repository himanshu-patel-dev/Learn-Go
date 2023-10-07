package main

import "fmt"

/*
Iterator is a behavioral design pattern that allows sequential traversal
through a complex data structure without exposing its internal details.
*/

func main() {
	user1 := &User{name: "first", age: 21}
	user2 := &User{name: "second", age: 25}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	userIterator := userCollection.createIterator()

	for userIterator.hasNext() {
		user := userIterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}

//User is &{name:first age:21}
//User is &{name:second age:25}
