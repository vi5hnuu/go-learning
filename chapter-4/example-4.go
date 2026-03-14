package main

import (
	"fmt"
)

func main() {
	fmt.Println(message());
}

func message() string{
	arr := [...]string{
		"Hello",
		"World",
		"Go",
		"Programming",
	};
	return fmt.Sprintf("%s %s %s %s", arr[0], arr[1], arr[2], arr[3])
}