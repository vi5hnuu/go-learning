package main

import (
	"fmt"
)

func main() {
	fmt.Println(dataType("hello"))
	fmt.Println(dataType(true))
	fmt.Println(dataType(3.14))
	fmt.Println(dataType(-5))
	fmt.Println(dataType(uint(42)))
}

func dataType(v interface{}) string {
	switch v.(type) {
	case string:
		return fmt.Sprintf("%v is string", v)
	case bool:
		return fmt.Sprintf("%v is bool", v)
	case float32,float64:
		return fmt.Sprintf("%v is float", v)
	case int:
		return fmt.Sprintf("%v is int", v)
	case int8,int32,int64:
		return fmt.Sprintf("%v is int", v)
	default:
		return fmt.Sprintf("%v is unknown", v)
	}
}