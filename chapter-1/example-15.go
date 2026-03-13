package main

import (
	"fmt"
)

func main() {
	count := 5
	count *= 5
	count += 5
	fmt.Println("Count : ", count)
	count++
	fmt.Println("Count : ", count)
	count--
	fmt.Println("Count : ", count)
	count -= 5
	fmt.Println("Count : ", count)
	count /= 5
	fmt.Println("Count : ", count)
	count = 26
	count /= 6
	fmt.Println("Count : ", count)
}
