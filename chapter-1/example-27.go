package main

import (
	"fmt"
)

// func main() {
// 	count:=5;
	
// 	if count>5{
// 		message := "Greater than 5"
// 	}else{
// 		message := "Less than or equal to 5"
// 	}
// 	fmt.Println(message)
// }

func main() {
	count:=5;

	var message string
	if count>5{
		message = "Greater than 5"
	}else{
		message = "Less than or equal to 5"
	}
	fmt.Println(message)
}
