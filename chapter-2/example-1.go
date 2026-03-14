package main

import (
	"fmt"
)

func main() {
	input:=5;
	if input & 1 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}