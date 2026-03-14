package main

import (
	"fmt"
)

func main() {
	for i:=1;i<=100;i++{
		switch {
			case i%3==0 && i%5==0:{
				fmt.Printf("%v FizzBuzz\n", i)
			}
			case i%3==0:
				fmt.Printf("%v Fizz\n", i)
			case i%5==0:
				fmt.Printf("%v Buzz\n", i)
			default:
				fmt.Printf("%v\n", i)
		}
	}
}