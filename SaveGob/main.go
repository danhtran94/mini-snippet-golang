package main

import (
    "fmt"
    "os"
    "encoding/gob"
)

type Person struct {
    Name Name
    Email []Email
}

type Email struct {
    Kind string
    Address string
}

type Name struct {
    Family string
    Personal string
}

func main()  {
    person := Person{
        Name: Name{Family: "Newmarch", Personal: "Jan"},
        Email: []Email{
            Email{Kind: "home", Address: "jan@new.name"},
        },
    }
    saveGob("person.gob", person)
}

func saveGob(fileName string, key interface{})  {
    outFile, err := os.Create(fileName)
    checkError(err)
    encoder := gob.NewEncoder(outFile)
    err = encoder.Encode(key)
    checkError(err)
    outFile.Close()
}

func checkError(err error)  {
    if err != nil {
        fmt.Println("Fatal error", err.Error())
        os.Exit(1)
    }
}