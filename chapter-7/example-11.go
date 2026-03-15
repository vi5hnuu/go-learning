package main

import (
    "fmt"
)

type record struct{
    key   string
    valueType string
    data interface{}
}

type person struct{
    name string
    age  int
    email string
}

type animal struct{
    name string;
    category string;
}

func main() {
    m:=make(map[string]any);
    a:=animal{name: "Lion", category: "Wild"}
    p:=person{name: "John", age: 30, email: "john@example.com"}
    m["person"]=p
    m["animal"]=a
    m["isMarried"]=true
    m["height"]=1.75

    rs:=[]record{}
    for k,v:=range m{
        r:=newRecord(k,v)
        rs=append(rs,r)
    }

    for _,r:=range rs{
        fmt.Printf("Key: %s, Type: %s, Value: %v\n", r.key, r.valueType, r.data)
    }
}


func newRecord(key string,i interface{}) record{
    r:=record{}
    r.key=key
    switch v:=i.(type){
    case string:
        r.valueType="string"
        r.data=v
    case int:
        r.valueType="int"
        r.data=v
    case bool:
        r.valueType="bool"
        r.data=v
    case person:
        r.valueType="person"
        r.data=v
    case animal:
        r.valueType="animal"
        r.data=v
    default:
        r.valueType="unknown"
        r.data=v
    }
    return r
}