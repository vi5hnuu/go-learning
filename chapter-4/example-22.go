package main

import (
	"fmt"
)

type name string;

type location struct{
	x int
	y int
}

type size struct{
	width int
	height int
}

type dot struct{
	name
	location
	size
}

func main() {
	dots:=getDots();
	for _, d := range dots {
		fmt.Println(d)
	}
}

func getDots() []dot{
	var dot1 dot;
	dot2:=dot{}
	dot2.name="A"
	dot2.x=5
	dot2.y=10
	dot2.width=100
	dot2.height=200

	dot3:=dot{
		name:"B",
		location:location{
			x:15,
			y:30,
		},
		size:size{
			width:150,
			height:300,
		},
	}

	dot4:=dot{};
	dot4.name="C"
	dot4.location.x=20
	dot4.location.y=40
	dot4.size.width=200
	dot4.size.height=400
	return []dot{dot1, dot2, dot3, dot4}
}