package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r := rand.Intn(100)
	fmt.Println(r);
	for{
		fmt.Println(r);
		if r%2==0 {
			r--;
			continue;
		} else {
			break;
		}
	}
}