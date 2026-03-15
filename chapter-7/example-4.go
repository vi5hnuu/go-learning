package main

import (
    "fmt"
)

type Speaker interface{
    Speak() string
}

type Cat struct{
    name string
}

func (c Cat) Speak() string {
    return fmt.Sprintf("Meow, my name is %s", c.name)
}

func (c Cat) String() string {
    return fmt.Sprintf("Cat: %s", c.name)
}

func main() {
    var s Speaker
    s = Cat{name: "Whiskers"}
    chapter(s)
    fmt.Println(s)
}

func chapter(s Speaker){
    fmt.Println(s.Speak())
}
