package main

import (
    "fmt"
)


func main() {
    calculator(add, 5, 10);
    calculator(substract, 5, 10);
}

func add(i,j int) string{
    result:=i+j;
    return fmt.Sprintf("Added %d + %d = %d", i, j, result);
}

func substract(i,j int) string{
    result:=i-j;
    return fmt.Sprintf("Substracted %d - %d = %d", i, j, result);
}


func calculator(f func(int, int) string,i,j int){
    fmt.Println(f(i,j));
}