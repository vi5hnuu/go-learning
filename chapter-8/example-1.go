package main

import (
    "fmt"
)

func main() {
    fmt.Println(findMaxInt([]int{1, 2, 3, 4, 5}))
    fmt.Println(findMaxFloat([]float64{1.1, 2.2, 3.3, 4.4, 5.5}))
}


func findMaxInt(nums []int) int{
    maxInt:=-1;
    if len(nums)<=0{
        return -1;
    }
    for i:=0;i<len(nums);i++{
        if nums[i]>maxInt{
            maxInt=nums[i];
        }
    }
    return maxInt;
}

func findMaxFloat(nums []float64) float64{
    maxFloat:=-1.0;
    if len(nums)<=0{
        return -1.0;
    }
    for i:=0;i<len(nums);i++{
        if nums[i]>maxFloat{
            maxFloat=nums[i];
        }
    }
    return maxFloat;
}