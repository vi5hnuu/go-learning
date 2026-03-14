package main

import (
	"fmt"
	"errors"
)

func main() {
	res,_:=doubler(42)
	fmt.Println(res)

	res,_=doubler("Hello")
	fmt.Println(res)

	res,e:=doubler(3.14)
	if e!=nil{
		fmt.Println(e)
	}else{
		fmt.Println(res)
	}
}

func doubler(v interface{}) (string,error){
	if i,ok:=v.(int);ok{
		return fmt.Sprintf("%d",i*2),nil
	}else if s,ok:=v.(string);ok{
		return s+s,nil
	}
	return "",errors.New("unsupported type")
}