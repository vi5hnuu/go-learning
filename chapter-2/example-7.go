package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Monday
	switch {
	case dayBorn == time.Saturday || dayBorn == time.Sunday:
		fmt.Println("Born on a weekend")
	default:
	fmt.Println("Born on a weekday")
	}
}