package main

type Iterator interface {
	hasNext() bool
	getNext() *User
}

// Concrete Implementation

type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) hasNext() bool {
	if u.index >= len(u.users) {
		return false
	}
	return true
}

func (u *UserIterator) getNext() *User {
	if !u.hasNext() {
		return nil
	}
	user := u.users[u.index]
	u.index += 1
	return user
}
