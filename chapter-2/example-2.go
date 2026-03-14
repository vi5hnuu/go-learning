package main

import (
	"fmt"
)

func main() {
	input:=-10;
	if input<0 {
		fmt.Println("Input can't be negative")
	}else if input & 1 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}