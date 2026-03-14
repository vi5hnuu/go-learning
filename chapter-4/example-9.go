package main

import (
	"fmt"
	"os"
)

func main() {
	locales:=getLocales(getPassedArgs(2))
	fmt.Println("Locales to use:",locales);
}

func getPassedArgs(minArgs int) []string{
	var args []string
	if len(os.Args) < minArgs {
		fmt.Printf("Not enough arguments. Expected at least %d, got %d\n", minArgs, len(os.Args))
		os.Exit(1)
	}
	for i:=1;i<len(os.Args);i++{
		args=append(args,os.Args[i])
	}
	return args
}

func getLocales(extraLocales []string) []string{
	var locales []string
	locales=append(locales,"en_US","fr_FR")
	locales=append(locales,extraLocales...)
	return locales
}