package main

import "fmt"

var name ="vishnu kumar"

func init(){
	fmt.Println("Initializing package...", name)
}

func init(){
	fmt.Println("Another initialization...")
}

func init(){
	fmt.Println("Third initialization...")
}

func main(){
	fmt.Println("Main function...")
}