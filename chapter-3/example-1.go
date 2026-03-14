package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(passwordChecked("Aa1!aaaa"))
	fmt.Println(passwordChecked("MaaPaa#"))

}

func passwordChecked(pw string) bool{
	pwR:=[]rune(pw)
	if len(pwR)<8{
		return false;
	}
	hasUpper,hasLower,hasNumber,hasSymbol:=false,false,false,false
	for _,c:=range pwR{
		switch {
		case unicode.IsUpper(c):
			hasUpper=true
		case unicode.IsLower(c):
			hasLower=true
		case unicode.IsNumber(c):
			hasNumber=true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSymbol=true
		}
	}
	return hasUpper&&hasLower&&hasNumber&&hasSymbol
}