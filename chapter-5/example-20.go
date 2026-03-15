package main

import (
    "fmt"
)

func main() {
    // example1();
    // example2();
    // example3();
    example4();
}

func example1(){
    for i := 1; i <= 5; i++ {
        defer func(i int){
            fmt.Printf("example-20: %d\n", i)
        }(i)
    }
}

func example2(){
    for i := 1; i <= 5; i++ {
        defer func(){
            fmt.Printf("example-20: %d\n", i)
        }()
    }
}

func example3(){
    for i := 1; i <= 5; i++ {
        func(){
            fmt.Printf("example-20: %d\n", i)
        }()
    }
}

func example4(){
    for i := 1; i <= 5; i++ {
        func(i int){
            fmt.Printf("example-20: %d\n", i)
        }(i)
    }
}