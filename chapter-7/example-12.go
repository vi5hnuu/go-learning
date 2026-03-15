package main

import (
    "fmt"
    "os"
    "errors"
)

type Payer interface{
    Pay() (string,float64)
}

type Developer struct{
    Individual Employee
    HourlyRate float64
    HoursWorkedInYear float64
    Review map[string]any
}

type Employee struct{
    Id int
    FirstName string
    LastName string
}

type Manager struct{
    Individual Employee
    Salary float64
    CommissionRate float64
}

func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"
	d := Developer{Individual: Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	m := Manager{Individual: Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}
	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	payDetails(d)
	payDetails(m)
}

func payDetails(p Payer){
    fullName, amount := p.Pay()
    fmt.Printf("Paying %s an amount of %.2f\n", fullName, amount);
}

func (d Developer) FullName() string {
    return fmt.Sprintf("%s %s", d.Individual.FirstName, d.Individual.LastName)
}

func (d Manager) FullName() string {
    return fmt.Sprintf("%s %s", d.Individual.FirstName, d.Individual.LastName)
}

func (d Developer) Pay() (string, float64) {
    fullName := d.FullName()
    amount := d.HourlyRate * d.HoursWorkedInYear
    return fullName, amount
}

func (d Developer) ReviewRating() error{
    total:=0;
    for _,v:=range d.Review{
        rating,err:=OverallReview(v);
        if err!=nil{
            return err
        }
        total+=rating;
    }
    averageRating:=float64(total)/float64(len(d.Review));
    fmt.Printf("Average Review Rating for %s: %.2f\n", d.FullName(), averageRating);
    return nil
}

func (m Manager) Pay() (string, float64) {
    fullName := m.FullName()
    amount := m.Salary + (m.CommissionRate * m.Salary)
    return fullName, amount
}

func OverallReview(i interface{}) (int,error){
    switch v := i.(type) {
    case int:
        return v,nil
    case string:
        rating,err:=convertReviewToInt(v);
        if err!=nil{
            return 0,err
        }
        return rating,nil
    default:
        return 0, errors.New("unknown type")
    }
}

func convertReviewToInt(str string) (int, error) {
	switch str {
	case "Excellent":
		return 5, nil
	case "Good":
		return 4, nil
	case "Fair":
		return 3, nil
	case "Poor":
		return 2, nil
	case "Unsatisfactory":
		return 1, nil
	default:
		return 0, errors.New("invalid rating: " + str)
	}
}