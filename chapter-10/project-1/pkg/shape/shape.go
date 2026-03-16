package shape

import "fmt"

type Shape interface {
    area() float64
    name() string
}

type Circle struct {
    Radius float64
}

type Square struct {
    Side float64
}

type Triangle struct {
    Base   float64
    Height float64
}

func (c Circle) name() string {
    return "Circle"
}

func (c Circle) area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (c Circle) String() string {
    return fmt.Sprintf("Shape: %s, Area: %.2f", c.name(), c.area())
}

func (s Square) name() string {
    return "Square"
}

func (s Square) area() float64 {
    return s.Side * s.Side
}

func (s Square) String() string {
    return fmt.Sprintf("Shape: %s, Area: %.2f", s.name(), s.area())
}

func (t Triangle) name() string {
    return "Triangle"
}

func (t Triangle) area() float64 {
    return 0.5 * t.Base * t.Height
}

func (t Triangle) String() string {
    return fmt.Sprintf("Shape: %s, Area: %.2f", t.name(), t.area())
}

func PrintShapeDetails(shapes ...Shape) {
    for _, shape := range shapes {
        fmt.Println(shape)
    }
}
