package main

import (
    "fmt"
)

type Employee struct{
    id int
    firstName string
    lastName string
}

type Developer struct{
    individual Employee
    hourlyRate int
    workWeek [7]int
}

type WeekDay int;
const (
    Sunday WeekDay = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func (d *Developer) HoursWorked(day WeekDay) int {
    return d.workWeek[day]
}

func main() {
    dev := Developer{
        individual: Employee{
            id:        1,
            firstName: "John",
            lastName:  "Doe",
        },
        hourlyRate: 50,
        workWeek:   [7]int{8, 8, 8, 8, 8, 0, 0},
    }

    fmt.Printf("Hours worked on Monday: %d\n", dev.HoursWorked(Monday))
}
