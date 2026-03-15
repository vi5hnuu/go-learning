package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if(len(os.Args)<3){
        fmt.Println("Please provide start,end as arguments.");
        os.Exit(1);
	}
    start,startError:=strconv.Atoi(os.Args[1]);
    end,endError:=strconv.Atoi(os.Args[2]);
    if startError != nil || endError != nil {
        fmt.Println("Invalid input. Please provide valid integers.");
        os.Exit(1);
    }else if start<0 || end<0 || end<start {
        fmt.Println("Invalid input. Please provide valid ranges.");
        os.Exit(1);
    }
    checkNNumbers(start, end);
}

func checkNNumbers(from,to int){
	fmt.Println("Checking numbers from", from, "to", to);
	for i:=from;i<=to;i++{
		if i&1==1{
			fmt.Printf("%d is odd\n", i)
		}else{
			fmt.Printf("%d is even\n", i)
		}
	}
	fmt.Println("Checking numbers ended");
}
