package main

import (
	"fmt"
	"net"
	"os"
	"encoding/gob"
)

type Person struct {
	Name Name
	Email []Email
}

type Name struct {
	Family string
	Personal string
}

type Email struct {
	Kind string
	Address string
}

func (p Person) String() string {
	s := p.Name.Personal +" "+ p.Name.Family
	for _, v := range p.Email {
		s += "\n"+ v.Kind +": "+ v.Address
	}
	return s
}


func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	fmt.Println("Listening port 1200 ...")
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)

		for n := 0; n < 10; n++ {
			person := new(Person)
			decoder.Decode(person)
			fmt.Println(person.String())
			encoder.Encode(*person)
		}
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal err:", err.Error())
		os.Exit(1)
	}
}
