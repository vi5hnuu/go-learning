package main

import (
    "fmt"
)

func main() {
    age:=20;
    name:="John"
    defer personAge(name,age)
    age+=10;
    defer personAge(name,age)
}

func personAge(name string,age int){
    fmt.Printf("%s is %d years old.\n", name, age)
}