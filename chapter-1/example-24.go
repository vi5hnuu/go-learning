package main

import (
	"fmt"
)

const (
	sunday=0
	monday=1
	tuesday=2
	wednesday=3
	thursday=4
	friday=5
	saturday=6
)

const (
	Sunday=iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday,sunday)
	fmt.Println(Monday,monday)
	fmt.Println(Tuesday,tuesday)
	fmt.Println(Wednesday,wednesday)
	fmt.Println(Thursday,thursday)
	fmt.Println(Friday,friday)
	fmt.Println(Saturday,saturday)
}
