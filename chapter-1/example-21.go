package main

import (
	"fmt"
)

func add5(count *int){
	*count+=5;
}

func main() {
	var count int
	add5(&count)
	fmt.Println(count);
}
