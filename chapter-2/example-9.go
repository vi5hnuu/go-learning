package main

import (
	"fmt"
)

func main() {

	names := []string{"vishnu kumar","krishan kumar","laxmi kumari","sita devi","geeta devi"}
	for i:=0;i<len(names);i++{
		fmt.Println(names[i])
	}

	for i, l := 0, len(names); i < l; i++ {
		fmt.Println(names[i])
	}
}