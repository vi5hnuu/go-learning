package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn:=time.Monday;
	switch dayBorn {
	case time.Sunday:
		fmt.Println("Born on a Sunday")
	case time.Monday:
		fmt.Println("Born on a Monday")
	case time.Tuesday:
		fmt.Println("Born on a Tuesday")
	case time.Wednesday:
		fmt.Println("Born on a Wednesday")
	case time.Thursday:
		fmt.Println("Born on a Thursday")
	case time.Friday:
		fmt.Println("Born on a Friday")
	case time.Saturday:
		fmt.Println("Born on a Saturday")
	default:
	fmt.Println("Invalid day")
	}
}