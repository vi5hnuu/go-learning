package main

import (
    "fmt"
)

func main() {
    fmt.Println("example-17 start")
    defer func(){
        fmt.Println("All tasks are done! --1")
    }()
    defer func(){
        fmt.Println("All tasks are done! --2")
    }()
    defer func(){
        fmt.Println("All tasks are done! --3")
    }()
    
    fmt.Println("example-17 done")
}