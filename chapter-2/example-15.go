package main

import (
	"fmt"
)

func main() {
	nums:= []int{5,10,15,20,25,33,41,77,88,99,45,47,53,67,84}

	for _,n:=range nums{
		fmt.Printf("%v ", n);
	}

	bubbleSort(nums);
	fmt.Println("\nSorted array:\n")
	for _,n:=range nums{
		fmt.Printf("%v ", n);
	}
}

func bubbleSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j >0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums
}