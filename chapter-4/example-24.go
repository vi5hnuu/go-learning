package main

import (
	"fmt"
	"math"
)

func main() {
	var i8 int8=math.MaxInt8
	m:=fmt.Sprintf("int8 =%v , int64 = %v\n",i8,int64(i8))
	fmt.Println(m)

	i:=128
	m += fmt.Sprintf("int = %v , int8 = %v\n", i, int8(i))
	fmt.Println(m)

	m += fmt.Sprintf("int8 = %v , float64 = %v\n", i8,float64(i8))
	fmt.Println(m)

	f64:=3.14
	m += fmt.Sprintf("float64 = %v , int = %v\n", f64, int(f64))
	fmt.Println(m)
}