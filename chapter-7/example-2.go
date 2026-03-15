package main

import (
    "fmt"
)

type Speaker interface{
    Speak() string
}

type cat struct{
    name string
}

func (c cat) Speak() string {
    return fmt.Sprintf("Meow, my name is %s", c.name)
}

func (c cat) String() string {
    return fmt.Sprintf("Cat: %s", c.name)
}

func main() {
    var s Speaker
    s = cat{name: "Whiskers"}
    fmt.Println(s.Speak())
    fmt.Println(s)
}
