package main

import (
    "fmt"
)

type Speaker interface {
    Speak() string
}

type cat struct{
    name string
}

func main() {
    c := cat{name: "Whiskers"}
    fmt.Println(c.Speak())
    emptyDetails(c)
    emptyDetails(10)
    emptyDetails("jane")
}

func (c cat) Speak() string {
    return fmt.Sprintf("%s says: Meow!", c.name)
}
func emptyDetails(i interface{}) {
    fmt.Printf("Empty details for:%v : %T\n", i,i)
}