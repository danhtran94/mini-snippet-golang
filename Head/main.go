package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := "http://www.golang.com"
	
	response, err := http.Head(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	
	fmt.Println(response.Status)
	for k, s := range  response.Header {
		fmt.Print(k +": ")
		for _, v := range s {
			fmt.Print(v +", ")
		}
		fmt.Println("")
	}
	
	os.Exit(0)
}

