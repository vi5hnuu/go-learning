package main

import (
    "fmt"
)

type Speaker interface{
    Speak() string
}

type Person struct{
    name      string
    age       int
    isMarried bool
}


func (p Person) Speak() string {
    return fmt.Sprintf("Hello, my name is %s, I am %d years old and it is %t that I am married.",
        p.name, p.age, p.isMarried);
}

func (p Person) String() string {
    return fmt.Sprintf("Person: %s, Age: %d, Married: %t",
        p.name, p.age, p.isMarried);
}

func main() {
    p := Person{name: "Alice", age: 30, isMarried: false}
    fmt.Println(p.Speak())
    fmt.Println(p)
}
