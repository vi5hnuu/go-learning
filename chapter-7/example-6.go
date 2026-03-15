package main

import (
    "fmt"
)

type Shape interface {
    Area() float64
    Name() string
}

type Circle struct {
    radius float64
}

type Square struct {
    side float64
}

type Triangle struct {
    base   float64
    height float64
}

func (c Circle) Name() string {
    return "Circle"
}

func (c Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}

func (c Circle) String() string {
    return fmt.Sprintf("Shape: %s, Area: %.2f", c.Name(), c.Area())
}

func (s Square) Name() string {
    return "Square"
}

func (s Square) Area() float64 {
    return s.side * s.side
}

func (s Square) String() string {
    return fmt.Sprintf("Shape: %s, Area: %.2f", s.Name(), s.Area())
}

func (t Triangle) Name() string {
    return "Triangle"
}

func (t Triangle) Area() float64 {
    return 0.5 * t.base * t.height
}

func (t Triangle) String() string {
    return fmt.Sprintf("Shape: %s, Area: %.2f", t.Name(), t.Area())
}

func main() {
    shapes := []Shape{
        Circle{radius: 5},
        Square{side: 4},
        Triangle{base: 3, height: 6},
    }

    printShapes(shapes...)
}

func printShapes(shapes ...Shape) {
    for _, shape := range shapes {
        fmt.Println(shape)
    }
}
