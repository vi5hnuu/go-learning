package main

import (
    "fmt"
)

// func main() {
//     fmt.Println("example-17 start")
//     done()
//     fmt.Println("example-17 done")
// }

func main() {
    fmt.Println("example-17 start")
    defer done()
    fmt.Println("example-17 done")
}

func done(){
    fmt.Println("All tasks are done!")
}