package main

import "fmt"

type user struct {
	name string
}

func createUser(name string) *user {
	return &user{
		name,
	}
}

type Iterator interface {
	HasNext() bool
	Next() *user
}

type iterator struct {
	index int
	users []*user
}

func (i *iterator) HasNext() bool {
	if i.index < len(i.users) {
		return true
	}
	return false
}

func (i *iterator) Next() *user {
	if i.HasNext() {
		user := i.users[i.index]
		i.index++
		return user
	}
	return nil
}

type Collection interface {
	CreateIterator() iterator
}

type collection struct {
	users []*user
}

func (c *collection) CreateIterator() Iterator {
	return &iterator{
		users: c.users,
	}
}

func main() {
	user1 := createUser("a")
	user2 := createUser("b")
	user3 := createUser("c")

	userCollection := &collection{
		users: []*user{user1, user2, user3},
	}

	userIterator := userCollection.CreateIterator()

	for userIterator.HasNext() {
		user := userIterator.Next()
		fmt.Printf("User is %s\n", user.name)
	}
}
