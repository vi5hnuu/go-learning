package main

import (
    "fmt"
    "errors"
)

var (
    ErrInvalidLastName   = errors.New("invalid last name")
    ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

type directDeposit struct{
    lastName string
    firstName string
    bankName string
    routingNumber int
    accountNumber int
}

func (d directDeposit) validateRoutingNumber() {
    defer func() {
        if r := recover(); r == ErrInvalidRoutingNumber {
            fmt.Println(r)
        }
    }()
    if d.routingNumber < 100 {
        panic(ErrInvalidRoutingNumber)
    }
}

func (d directDeposit) validateLastName() {
    defer func() {
        if r := recover(); r == ErrInvalidLastName {
            fmt.Println(r)
        }
    }()
    if d.lastName == "" {
        panic(ErrInvalidLastName)
    }
}

func (d directDeposit) report() {
    d.validateLastName()
    d.validateRoutingNumber()
    fmt.Println(d)
}

func main() {
    dd := directDeposit{
        lastName:     "Doe",
        firstName:    "John",
        bankName:     "Bank of Go",
        routingNumber: 123456789,
        accountNumber: 987654321,
    }
    dd.report()

    dm := directDeposit{
        lastName:     "Smith",
        firstName:    "Jane",
        bankName:     "Bank of Go",
        routingNumber: 10,
        accountNumber: 40,
    }
    dm.report()
}
