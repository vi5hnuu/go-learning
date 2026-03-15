package main

import (
    "fmt"
)

type Speaker interface{
    Speak() string
}

type cat struct{}

func (c cat) Speak() string {
    return "Meow"
}

func main() {
    var s Speaker
    s = cat{}
    fmt.Println(s.Speak())
}
