package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn:=time.Monday;
	switch dayBorn {
	case time.Monday:fallthrough
	case time.Tuesday:fallthrough
	case time.Wednesday:fallthrough
	case time.Thursday:fallthrough
	case time.Friday:
		fmt.Println("Born on a weekday")
	case time.Saturday:
	case time.Sunday:
		fmt.Println("Born on a weekend")
	default:
	fmt.Println("Invalid day")
	}
}