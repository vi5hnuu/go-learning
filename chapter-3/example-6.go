package main

import (
	"fmt"
)

func main() {
	comment1:=`This is the Best
	thing ever!`
	comment2:=`This is the Worst\nthing ever!`
	comment3:="This is the Best\nthing ever!"

	fmt.Println(comment1)
	fmt.Println(comment2)
	fmt.Println(comment3)
}