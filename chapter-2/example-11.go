package main

import (
	"fmt"
)

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You": 3,
		"Give": 2,
		"Never": 1,
		"Up": 4,
	}

	var maxCountWord string
	var maxCount int
	for key, value := range words {
		if value>maxCount{
			maxCountWord,maxCount=key,value;
		}
	}
	fmt.Printf("Word with maximum count: %s (%d)\n", maxCountWord, maxCount)
}