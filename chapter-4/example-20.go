package main

import (
	"fmt"
)

type user struct{
	name string
	age int
	balance float64
	member bool
}

func main() {
	users:=getUsers()
	fmt.Println(users)
}


func getUsers() []user {
	users:=[]user{
		{"Alice", 30, 1000.50, true},
		{"Bob", 25, 500.75, false},
		{name:"Charlie", age:35, balance:2000.00, member:true},
	}

	user0:=user{
		name: "Charlie",
		age: 35,
		balance: 2000.00,
		member: true,
	}
	users = append(users, user0)

	user1:=user{
		name: "Charlie",
		balance: 2000.00,
		member: true,
	}
	user1.age=55;
	users = append(users, user1)
	return users
}