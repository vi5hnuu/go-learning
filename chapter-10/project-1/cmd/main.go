package main

import (
	"project1/pkg/shape"
)

func main() {
    circle := shape.Circle{Radius: 5}
    square := shape.Square{Side: 4}
    triangle := shape.Triangle{Base: 3, Height: 6}

    shape.PrintShapeDetails(circle, square, triangle)
}
