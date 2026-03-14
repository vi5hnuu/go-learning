package main

import (
	"fmt"
)

func main() {
	words:=[4]string{"Hello", "World", "Go", "Programming"};
	words[0]="Hi";
	fmt.Println(fmt.Sprintf("%s %s %s %s", words[0], words[1], words[2], words[3]));
}