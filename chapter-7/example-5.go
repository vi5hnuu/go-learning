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

type Dog struct{
    name string
}

type Horse struct{
    name string
}


func (c Cat) Speak() string {
    return fmt.Sprintf("Meow, my name is %s", c.name)
}

func (c Cat) String() string {
    return fmt.Sprintf("Cat: %s", c.name)
}

func (d Dog) Speak() string {
    return fmt.Sprintf("Woof, my name is %s", d.name)
}

func (d Dog) String() string {
    return fmt.Sprintf("Dog: %s", d.name)
}

func (h Horse) Speak() string {
    return fmt.Sprintf("Neigh, my name is %s", h.name)
}

func (h Horse) String() string {
    return fmt.Sprintf("Horse: %s", h.name)
}

func main() {
    var s Speaker
    s = Cat{name: "Whiskers"}
    chapter(s)
    fmt.Println(s)
    s = Dog{name: "Rex"}
    chapter(s)
    fmt.Println(s)
    s = Horse{name: "Spirit"}
    chapter(s)
    fmt.Println(s)
}

func chapter(s Speaker){
    fmt.Println(s.Speak())
}
