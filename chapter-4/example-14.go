package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid user id");
		os.Exit(1)
	}
	userID := os.Args[1]

	users:=getUsers();
	name,ok:=users[userID];

	if !ok{
		for id, user := range users {
			fmt.Printf("ID: %s, Name: %s\n", id, user)
		}
		os.Exit(1)
	}
	fmt.Printf("User: %s\n", name)
}

func getUsers() map[string]string{
	users := map[string]string{
		"user1": "John Doe",
		"user2": "Jane Smith",
	}
	users["user3"] = "Alice Johnson"
	return users
}