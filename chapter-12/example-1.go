package main

import (
	"fmt"
	"time"
)

func main() {
    start := time.Now()
    fmt.Println("The script started at ", start)
    fmt.Println("saving...")
    time.Sleep(2 * time.Second)
    end := time.Now()
    fmt.Println("The script ended at ", end)
    fmt.Println("Total time taken: ", end.Sub(start))
}
