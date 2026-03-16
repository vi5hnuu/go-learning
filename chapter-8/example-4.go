package main

import (
    "fmt"
)

func main() {
    FindLargestRanchStock(map[string]float64{
        "Ranch1": 100.5,
        "Ranch2": 200.75,
        "Ranch3": 150.0,
    })
}

func FindLargestRanchStock[K comparable, V int | float64](m map[K]V) K {
    var stock V
    var name K
    for k, v := range m {
        if v > stock {
            stock = v
            name = k
        }
    }
    fmt.Printf("Largest ranch stock is %v with %v\n", name, stock)
    return name
}