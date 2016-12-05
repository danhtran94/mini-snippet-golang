package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"net/http/httputil"
	"strings"
)

func main() {
	url := "http://www.golang.com"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error: "+ err.Error())
		os.Exit(1)
	}
	
	if res.Status != "200 OK" {
		fmt.Println(res.Status)
		os.Exit(2)
	}
	
	b, _ := httputil.DumpResponse(res, false)
	fmt.Println(string(b))
	
	contentTypes := res.Header["Content-Type"]
	if !acceptableCharset(contentTypes) {
		fmt.Println("Cannot handle", contentTypes)
		os.Exit(4)
	}
	
	var buf [512]byte
	reader := res.Body
	for {
		n, err := reader.Read(buf[0:])
		fmt.Print(string(buf[0:n]))
		if err != nil {
			os.Exit(0)
		}
	}
	os.Exit(0)
}

func acceptableCharset(contentTypes []string) bool {
	for _, cType := range contentTypes {
		if strings.Index(cType, "utf-8") != -1 {
			return true
		}
	}
	return false
}