package main

type Collection interface {
	createIterator() Iterator
}

// Concrete Implementation

type UserCollection struct {
	users []*User
}

func (u *UserCollection) createIterator() Iterator {
	return &UserIterator{
		index: 0,
		users: u.users,
	}
}
