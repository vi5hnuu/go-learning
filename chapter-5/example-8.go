package main

import (
    "fmt"
)

func main() {
    fmt.Println(addNums(1, 2, 3))
    fmt.Println(addNums(4, 5, 6, 7))
    fmt.Print(addNums())
}


func addNums(i ...int) int {
    fmt.Println(i)
    fmt.Printf("%T\n", i)
    fmt.Printf("Len: %d\n", len(i))
    fmt.Printf("Cap: %d\n", cap(i))

    sum := 0
    for _, v := range i {
        sum += v
    }
    return sum
}