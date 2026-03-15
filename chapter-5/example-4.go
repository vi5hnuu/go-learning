package main

import (
    "fmt"
    "strings"
)

func main() {
    headers1:=[]string{"empid","employee","address","hours worked","hourly wage","manager"}
    csvHeaderCol(headers1);

    headers2:=[]string{"employee","empId","hours worked","address","manager","hourly wage"}
    csvHeaderCol(headers2);
}

func csvHeaderCol(headers []string) {
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
    fmt.Println(csvHeadersToColumnIndex)
}
