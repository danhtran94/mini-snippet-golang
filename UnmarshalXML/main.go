package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"io/ioutil"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name Name `xml:"name"`
	Email []Email `xml:"email"`
}

type Name struct {
	Family string `xml:"family"`
	Personal string `xml:"personal"`
}

type Email struct {
	Type string `xml:"type,attr"`
	Address string `xml:",chardata"`
}

func main() {
	xmlBytes, err := ioutil.ReadFile("person.xml")
	if err != nil {
		log.Fatal("Fatal error:", err.Error())
		os.Exit(1)
	}
	person := new(Person)
	
	err = xml.Unmarshal(xmlBytes, person)
	/*
	decoder := xml.NewDecoder(bytes.NewBuffer(xmlBytes))
	decoder.Strict = false
	decoder.AutoClose = xml.HTMLAutoClose
	decoder.Entity = xml.HTMLEntity
	decoder.Decode(person)
	*/
	
	if err != nil {
		log.Fatal("Fatal error:", err.Error())
		os.Exit(1)
	}
	
	fmt.Println("Root XML tagName: "+ person.XMLName.Local)
	fmt.Println("Family name: \""+ person.Name.Family +"\"")
	fmt.Println("First email: \""+ person.Email[0].Address +"\"")
	
	fmt.Println("Marshal ...")
	bytes, err := xml.Marshal(person)
	fmt.Println(string(bytes))
}

