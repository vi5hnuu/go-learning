package main

import (
	"fmt"
)

func main() {
		fmt.Println("Users", getUsers())
}

func getUsers() map[string]string{
	users := map[string]string{
		"user1": "John Doe",
		"user2": "Jane Smith",
	}
	users["user3"] = "Alice Johnson"
	return users
}
