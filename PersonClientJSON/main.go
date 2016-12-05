package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"os"
	"net-programing/xmlcodec"
)

type Person struct {
	Name string
	Emails []string
}

func main() {
	conn, err := websocket.Dial("ws://localhost:12345/", "", "http://localhost")
	checkError(err)
	
	person := Person{
		Name: "Jan",
		Emails: []string{"abc@abc", "xyz@xyz"},
	}
	
	// err = websocket.JSON.Send(conn, person)
	err = xmlcodec.XMLCodec.Send(conn, person)
	if err != nil {
		fmt.Println("Couldn't send msg "+ err.Error())
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
