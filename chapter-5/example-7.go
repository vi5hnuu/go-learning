package main

import (
    "fmt"
    "strings"
)

func main() {
    headers1:=[]string{"empid","employee","address","hours worked","hourly wage","manager"}
    r1:=csvHeaderCol(headers1);
    fmt.Println(r1)

    headers2:=[]string{"employee","empId","hours worked","address","manager","hourly wage"}
    r2:=csvHeaderCol(headers2);
    fmt.Println(r2)
}

func csvHeaderCol(headers []string) map[int]string {
    csvHeadersToColumnIndex:=make(map[int]string);

    for i, header := range headers {
        header=strings.TrimSpace(header)
        switch strings.ToLower(header) {
        case "employee":
            csvHeadersToColumnIndex[i] = "employee"
        case "hours worked":
            csvHeadersToColumnIndex[i] = "hours worked"
        case "hourly wage":
            csvHeadersToColumnIndex[i] = "hourly wage"
        }
    }
    return csvHeadersToColumnIndex
}
