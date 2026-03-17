package main

import (
	"fmt"
	"time"
)

func main() {
    t1, err := time.Parse(time.RFC3339, "2019-09-27T22:18:11+00:00")
    if err != nil {
        fmt.Println(err)
    }
    t2, err := time.Parse(time.UnixDate, "2019-09-27T22:18:11+00:00")
    if err != nil {
        fmt.Println(err)
    }
    t3, err := time.Parse(time.ANSIC, "2019-09-27T22:18:11+00:00")
    if err != nil {
    fmt.Println(err)
}
    fmt.Println("RFC3339:", t1)
    fmt.Println("UnixDate", t2)
    fmt.Println("ANSIC", t3)
}
