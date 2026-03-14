package main

import (
	"fmt"
)

func main() {
	daysOfWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday","Sunday"}
	daysOfWeek=append(daysOfWeek[len(daysOfWeek)-1:],daysOfWeek[:len(daysOfWeek)-1]...);
	fmt.Println(daysOfWeek)
}