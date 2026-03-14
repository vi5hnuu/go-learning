package main

import (
	"fmt"
)

func main() {
	words:=[]string{
		"Good",
		"Good",
		"Bad",
		"Good",
		"Good",
	}

	fmt.Println(append(append([]string{},words[0:2]...),words[3:]...))
}