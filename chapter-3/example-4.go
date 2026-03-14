package main

import (
	"fmt"
)

func main() {
	var num1 uint8 = 1<<8-1
	// var num2 uint8 = 1<<8

	fmt.Println(num1)
	// fmt.Println(num2)
	
	num1++;
	fmt.Println(num1)
}