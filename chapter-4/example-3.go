package main

import (
	"fmt"
)

func main() {
	var arr1 [10]int;
	arr2:=[...]int{9:0}
	arr3:=[10]int{1,9:10,4:5}

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(arr1 == arr2, arr1 == arr3)
}