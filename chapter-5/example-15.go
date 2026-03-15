package main

import (
    "fmt"
)

func main() {
    sq := square(5);
    fmt.Println(sq());
    fmt.Printf("Type of sq : %T\n", sq)
}


func square(x int) func() int {
    return func() int {
        return x * x
    }
}