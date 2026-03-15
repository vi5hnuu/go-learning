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

func (d directDeposit) validateRoutingNumber() error {
    if d.routingNumber < 100 {
        return ErrInvalidRoutingNumber
    }
    return nil
}

func (d directDeposit) validateLastName() error {
    if d.lastName == "" {
        return ErrInvalidLastName
    }
    return nil
}

func (d directDeposit) report() {
    if err := d.validateLastName(); err != nil {
        fmt.Println(err)
    }
    if err := d.validateRoutingNumber(); err != nil {
        fmt.Println(err)
    }
    fmt.Println("All validations passed")
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
}
