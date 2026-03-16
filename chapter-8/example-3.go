package main

import (
    "fmt"
)

type Number interface{
    int | float64
}

func main() {
    fmt.Println(findMax([]int{1, 2, 3, 4, 5}))
    fmt.Println(findMax([]float64{1.1, 2.2, 3.3, 4.4, 5.5}))
}


func findMax[Num Number](nums []Num) Num {
    if len(nums) <= 0 {
        var zero Num
        return zero
    }
    maxVal := nums[0];

    for i := 0; i < len(nums); i++ {
        if nums[i] > maxVal {
            maxVal = nums[i]
        }
    }
    return maxVal
}