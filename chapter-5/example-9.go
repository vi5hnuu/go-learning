package main

import (
    "fmt"
)

func main() {
    nums:=[]int{1, 2, 3, 4, 5}
    fmt.Println(addNums(nums...))
    fmt.Println(addNums())
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
