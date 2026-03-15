package main

import (
    "fmt"
)

func main() {
    itemsSold()
}


func itemsSold(){
    items:=make(map[string]int)
    items["john"]=5
    items["jane"]=50
    items["jow"]=40
    items["doe"]=30

    for person,quantity:=range items{
        fmt.Printf("%s sold %d quantities and", person, quantity)
        if quantity<40{
            fmt.Println(" is below expectations.")
        }else if quantity>=40 && quantity<=100{
            fmt.Println(" meets expectations.")
        }else{
            fmt.Println(" exceeds expectations.")
        }
    }
}