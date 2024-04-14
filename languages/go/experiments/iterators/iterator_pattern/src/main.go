package main

import "fmt"

type User struct {
	name string
	age  int
}

type iterator interface {
	hasNext() bool
	next()
	getCurrentItem() *User
}

type iterable interface {
	getIterator() iterator
}

type userIterator struct {
	index int
	users []*User
}

func (user *userIterator) hasNext() bool {
	return user.index < len(user.users)
}

func (user *userIterator) next() {
	if user.hasNext() {
		user.index += 1
	}
}

func (user *userIterator) getCurrentItem() *User {
	if user.hasNext() {
		return user.users[user.index]
	}

	return nil
}

type userIterableCollection struct {
	users []*User
}

func (user *userIterableCollection) getIterator() iterator {
	return &userIterator{users: user.users}
}

func main() {
	user1 := &User{name: "Frik", age: 22}
	user2 := &User{name: "Sam", age: 21}
	user3 := &User{name: "Claire", age: 24}
	user4 := &User{name: "Bella", age: 20}

	println("\nUser1 is:", user1.name, "( age:", user1.age, ")")

	println("\n\nLet's look at the entire list:")

	users := &userIterableCollection{users: []*User{user1, user2, user3, user4}}
	iterator := users.getIterator()

	for iterator.hasNext() {
		element := iterator.getCurrentItem()
		fmt.Printf("We have %s, aged %d.\n", element.name, element.age)
		iterator.next()
	}
}
