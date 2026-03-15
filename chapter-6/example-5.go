package main

import (
    "fmt"
)

func main() {
    a();
    fmt.Println("This line will now get printed from main() function")
}

func a(){
    b();
    fmt.Println("This line will now get printed from a() function")
}

func b(){
    defer func(){
        if r:=recover();r!=nil{
            fmt.Println("Recovered in b():", r)
        }
    }()
    panic("panic in b()")
}
