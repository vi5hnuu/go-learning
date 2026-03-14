package main

import (
	"fmt"
)

type point struct{
	x int
	y int
}

func main() {
	fmt.Println(compare())
}

func compare() (bool,bool){
	point1:=point{10,20}
	point2:=struct{
		x int
		y int
	}{
		x: 10,
		y: 20,
	}
	point3:=struct{
		x int
		y int
	}{
	}
	point3.x = 10
	point3.y = 20
	return point1 == point2, point1 == point3
}