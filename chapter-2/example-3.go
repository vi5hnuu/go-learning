package main

import (
	"fmt"
	"errors"
)

func main() {
	input:=21;
	if err:=validate(input);err!=nil{
		fmt.Println(err)
	}else if input & 1 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}

func validate(val int) error{
	if val < 0 {
		return errors.New("Input can't be negative")
	}else if val > 100 {
		return errors.New("Input can't be greater than 100")
	}else if val%7==0 {
		return errors.New("Input can't be a multiple/divisible of 7")
	}
	return nil
}