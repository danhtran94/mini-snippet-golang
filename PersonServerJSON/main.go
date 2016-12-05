package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"net/http"
	"os"
	"net-programing/xmlcodec"
)

type Person struct {
	Name string
	Emails []string
}

func ReceivePerson(ws *websocket.Conn) {
	var person Person
	// err := websocket.JSON.Receive(ws, &person)
	err := xmlcodec.XMLCodec.Receive(ws, &person)
	if err != nil {
		fmt.Println("Can't receive")
	} else {
		fmt.Println("Name: "+ person.Name)
		for _, e := range person.Emails {
			fmt.Println("An email: "+ e)
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(ReceivePerson))
	err := http.ListenAndServe(":12345", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

