package main

import "fmt"

func main() {

	user1 := &User{
		name: "Asliddin",
		age:  21,
	}

	user2 := &User{
		name: "Tursinbek",
		age:  19,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
