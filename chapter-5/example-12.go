package main

import (
    "fmt"
)

func main() {
    c1 := counter();
    c2 := counter();
    fmt.Println(c1())
    fmt.Println(c2())
    fmt.Println(c1())
}


func counter() func()int{
    count:=0;
    return func()int{
        count++;
        return count;
    }
}
