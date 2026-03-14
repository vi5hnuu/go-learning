package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(getPassedArgs(2))
}

func getPassedArgs(minArgs int) []string{
	if len(os.Args) < minArgs {
		fmt.Printf("Not enough arguments. Expected at least %d, got %d\n", minArgs, len(os.Args))
		os.Exit(1)
	}
	var args []string
	for i:=1;i<len(os.Args);i++{
		args=append(args,os.Args[i])
	}
	return args
}