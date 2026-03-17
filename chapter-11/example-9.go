package main

import (
	"errors"
	"log"
)

func main() {
    log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
    log.Println("Start of example-9")
    err := errors.New("Application aborted")
    if err != nil {
        log.Fatalln(err)
    }
    log.Println("End of example-9")
}
