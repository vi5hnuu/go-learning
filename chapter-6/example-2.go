package main

import (
    "fmt"
    "errors"
)

func main() {
    msg:="good bye"
    message(msg);
    fmt.Println("This line will not get printed");
}

func message(msg string){
    if msg=="good bye"{
        panic(errors.New("Something went wrong"));
    }
}
