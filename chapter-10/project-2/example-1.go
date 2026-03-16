package main

import "fmt"

var name ="vishnu kumar"

func init(){
	fmt.Println("Initializing package...", name)
}

func main(){
	fmt.Println("Name:", name)
}