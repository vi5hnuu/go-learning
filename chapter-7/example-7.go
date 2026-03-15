package main

import (
    "fmt"
    "encoding/json"
    "io"
    "strings"
)

type Person struct {
    Name    string
    Age     int
}

func (p Person) String() string {
    return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func main() {
    s:=`{"Name": "John", "Age": 30}`
    s2:=`{"Name": "Jane", "Age": 25}`
    p,err:=loadPerson(strings.NewReader(s));
    if err!=nil{
        fmt.Println("Error loading person:", err)
        return
    }
    fmt.Println("Loaded person:", p)

    p2,err2:=loadPerson2(s2)
    if err2!=nil{
        fmt.Println("Error loading person:", err2)
        return
    }
    fmt.Println("Loaded person:", p2)
}

func loadPerson2(s string)(Person ,error){
    var p Person
    err:=json.NewDecoder(strings.NewReader(s)).Decode(&p)
    if err != nil {
        return Person{}, err
    }
    return p,nil
}

func loadPerson(r io.Reader) (Person,error){
    var p Person
    err := json.NewDecoder(r).Decode(&p)
    if err != nil {
        return Person{}, err
    }
    return p,nil
}