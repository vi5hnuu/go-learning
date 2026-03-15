package main

import (
    "fmt"
)

func main() {
    f:=func(){
        fmt.Println("Hello from the anonymous function!")
    }
    fmt.Println("Line after the anonymous function!")
    f();
}
