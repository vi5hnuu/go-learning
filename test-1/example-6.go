package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	var seed int64 = 123456789
	rand.NewSource(seed)
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
}
