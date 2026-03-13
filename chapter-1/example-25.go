package main

import (
	"fmt"
)

var level = "pkg"
func main() {
	fmt.Println("Main start : ",level);
	if true{
		fmt.Println("Block start : ",level)
		funcA();
	}
}

func funcA() {
	fmt.Println("FuncA start : ",level)
	if true {
		fmt.Println("Block start : ",level)
		funcB()
	}
}

func funcB() {
	fmt.Println("FuncB start : ",level)
}
