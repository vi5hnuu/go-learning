package main

import (
    "fmt"
    "errors"
)

var (
    ErrInvalidLastName   = errors.New("invalid last name")
    ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

func main() {
    fmt.Println(ErrInvalidLastName)
    fmt.Println(ErrInvalidRoutingNumber)
}
