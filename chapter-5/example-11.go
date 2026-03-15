package main

import (
    "fmt"
)

func main() {
    j:=9;
    x:=func(i int) int{
        return i*i;
    }

    fmt.Println("The square of", j, "is", x(j))
}