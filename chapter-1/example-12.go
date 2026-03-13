package main

import (
	"fmt"
)

func main() {
	query, limit, offset := "bar", 10, 0

	fmt.Println(query, limit, offset)

	query, limit, offset = "foo", 20, 0
	fmt.Println(query, limit, offset)
}
