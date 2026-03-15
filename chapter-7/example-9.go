package main

import (
    "fmt"
)

func main() {
    var str interface{} = "some string";
    var i any = 42
    var b interface{} = true;
    fmt.Printf("%v -> %T, %v -> %T, %v -> %T\n", str,str, i,i, b,b)
}
