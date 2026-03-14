package main

import (
	"fmt"
)

func main() {
	bioData := map[string]string{
		"name":    "Vishnu Kumar",
		"age":     "26",
		"address": "Rajasthan",
	}

	for key,value := range bioData{
		fmt.Printf("%s: %s\n", key, value)
	}
}