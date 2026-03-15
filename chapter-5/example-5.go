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
    for i:=start;i<=end;i++{
		num,kind:=checkNNumbers(i);
		fmt.Printf("%d is %s\n", num, kind);
	}
}

func checkNNumbers(num int) (int,string){
	switch{
        case num&1==1:
            return num,"odd"
        default:
            return num,"even"
	}
}
