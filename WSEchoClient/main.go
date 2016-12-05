package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"io"
	"os"
)

func main()  {
	conn, err := websocket.Dial("ws://localhost:12345/", "", "http://localhost")
	checkError(err)
	var msg string
	for {
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Couldn't receive msg "+ err.Error())
			break
		}
		fmt.Println("Received from server: "+ msg)
		err = websocket.Message.Send(conn, msg)
		if err != nil {
			fmt.Println("Couldn't return msg")
			break
		}
	}
	os.Exit(0)
}

func checkError(err error)  {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}