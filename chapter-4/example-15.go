package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args)<2{
		fmt.Println("Invalid user id");
		os.Exit(1)
	}
	
	users:=map[string]string{
		"305":"Sue",
		"204":"Bob",
		"631":"Jake",
		"073":"Tracy",
	}

	userId:=os.Args[1]
	name,ok:=users[userId];
	if !ok{
		fmt.Println("User not found");
		os.Exit(1)
	}
	fmt.Printf("User: %s\n", name)
}