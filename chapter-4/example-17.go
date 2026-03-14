package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid user id")
		os.Exit(1)
	}
	userId := os.Args[1]
	users := getUsers()
	printUsers(users)
	deleteUser(users, userId)
	printUsers(users)
}

func deleteUser(users map[string]string, userId string) {
	delete(users, userId)
}

func printUsers(users map[string]string) {
	for id, name := range users {
		fmt.Printf("User ID: %s, Name: %s\n", id, name)
	}
}

func getUsers() map[string]string{
	users := map[string]string{
		"user1": "John Doe",
		"user2": "Jane Smith",
	}
	users["user3"] = "Alice Johnson"
	return users
}