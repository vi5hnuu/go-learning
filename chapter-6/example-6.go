package main

import (
    "fmt"
    "errors"
)

var ErrHourlyRate=errors.New("invalid hourly rate")
var ErrHoursWorked=errors.New("invalid hours worked per week")

func main() {
    pay := payDay(100, 25)
    fmt.Println(pay)
    pay = payDay(100, 200)
    fmt.Println(pay)
    pay = payDay(60, 25)
    fmt.Println(pay)
}

func payDay(hoursWorked,hourlyrate int) int{
    defer func(){
        if r:=recover();r!=nil{
            if r==ErrHourlyRate{
                fmt.Println("Recovered in payDay():", r)
            }else if r==ErrHoursWorked{
                fmt.Println("Recovered in payDay():", r)
            }
        }
    }()
    if hourlyrate<10 || hourlyrate>75{
        panic(ErrHourlyRate)
    }else if hoursWorked<0 || hoursWorked>80{
        panic(ErrHoursWorked)
    }else if hoursWorked > 40{
        hoursOver:=hoursWorked-40;
        overTime:=hoursOver*2;
        regularPay:=hoursWorked*hourlyrate;
        return regularPay+overTime
    }
    return hoursWorked*hourlyrate
}
