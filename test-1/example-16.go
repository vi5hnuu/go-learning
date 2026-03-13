package main

import (
	"fmt"
)

func main() {
	visits := 15
	fmt.Println("First visit : ", visits == 1)
	fmt.Println("Return visit : ", visits != 1)
	fmt.Println("Siver memebr : ", visits >= 10 && visits < 21)
	fmt.Println("Gold memebr : ", visits > 20 && visits <= 30)
	fmt.Println("Platinum memebr : ", visits > 30)
}
