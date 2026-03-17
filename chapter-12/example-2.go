package main

import (
	"fmt"
	"time"
)

func whatstheclock() string {
    return time.Now().Format(time.ANSIC)
}

func main(){
    now := time.Now()
    onlyAfter, err := time.Parse(time.RFC3339,"2027-11-01T22:08:41+00:00")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(now, onlyAfter)
    fmt.Println(now.After(onlyAfter))
    if now.After(onlyAfter) {
        fmt.Println("Executing actions!")
    } else {
        fmt.Println("Now is not the time yet!!")
    }
}