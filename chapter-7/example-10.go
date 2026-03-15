package main

import (
    "fmt"
    "strconv"
)

func main() {
    var str interface{} = "some string";
    var i any = 42
    var b interface{} = true;

    applyDefaults(str)
    applyDefaults(i)
    applyDefaults(b)
}

func applyDefaults(value interface{}) {
    switch v := value.(type) {
    case string:
        fmt.Printf("Applying default for string: %s\n", v)
    case int:
        if newValue, err := strconv.Atoi(fmt.Sprintf("%v", v)); err == nil {
            fmt.Printf("Applying default for int: %d\n", newValue+10)
        }else{
            fmt.Printf("Error converting int: %v\n", err)
        }
    case bool:
        fmt.Printf("Applying default for bool: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
