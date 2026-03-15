package main

import (
    "fmt"
)

func main() {
    devSalary:=salary(50,2080,developerSalary);
    bossSalary:=salary(60,2080,managerSalary);
    fmt.Println("Developer Salary: ", devSalary);
    fmt.Println("Boss Salary: ", bossSalary);
}

func salary(x,y int,f func(int,int) int) int{
    pay:=f(x,y);
    return pay;
}

func managerSalary(baseSalary int, bonus int) int {
    return baseSalary + bonus;
}

func developerSalary(hourlyRate int, hoursWorked int) int {
    return hourlyRate * hoursWorked;
}