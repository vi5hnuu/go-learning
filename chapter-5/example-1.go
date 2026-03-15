package main

import (
    "fmt"
)

func main() {
	fmt.Println("Main function in control");
	checkNNumbers();
	fmt.Println("Main function ended");
}

func checkNNumbers(){
	fmt.Println("Checking numbers from 1 to 30");
	for i:=1;i<=30;i++{
		if i&1==1{
			fmt.Printf("%d is odd\n", i)
		}else{
			fmt.Printf("%d is even\n", i)
		}
	}
	fmt.Println("Checking numbers ended");
}
