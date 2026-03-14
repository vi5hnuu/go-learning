package main

import (
	"fmt"
)

func main() {
	logLevel:="INFO"
	for index,runeVal:=range logLevel{
		fmt.Printf("Character at index %d is %c\n", index, runeVal)
	}
}