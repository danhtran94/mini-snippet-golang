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
	person := Person{
		Name: Name{Family: "Tran", Personal: "Danh"},
		Email: []Email{
			Email{Kind: "home", Address: "dd2672015@gmail.com"},
		},
	}

	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "host:port")
		os.Exit(1)
	}

	service := os.Args[1]
	remoteAddr, _ := net.ResolveTCPAddr("tcp4", service)
	conn, err := net.DialTCP("tcp", nil, remoteAddr)
	checkError(err)

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	for n := 0; n < 10; n++ {
		encoder.Encode(person)
		toNewPerson := new(Person)
		decoder.Decode(toNewPerson)
		fmt.Println(toNewPerson.String())
	}

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal err:", err.Error())
		os.Exit(1)
	}
}
