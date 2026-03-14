package main

import (
	"fmt"
)

func main() {
	s1,s2,s3:=getSlices();
	fmt.Println("slice 1 :",len(s1),cap(s1))
	fmt.Println("slice 2 :",len(s2),cap(s2))
	fmt.Println("slice 3 :",len(s3),cap(s3))
}

func getSlices()([]int,[]int,[]int){
	var s1 []int;
	s2:=make([]int,10)
	s3:=make([]int,10,50);
	return s1,s2,s3;
}