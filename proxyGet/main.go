package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"encoding/base64"
)

// const auth = "danhtran:password"

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "http://proxyhost:port http://host:port/page")
		os.Exit(1)
	}
	
	proxyString := os.Args[1]
	proxyURL, err := url.Parse(proxyString)
	checkError(err)
	
	rawURL := os.Args[2]
	url, err := url.Parse(rawURL)
	checkError(err)
	
	//basic := "Basic "+ base64.StdEncoding.EncodeToString([]byte(auth))
	
	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{Transport: transport}
	
	request, err := http.NewRequest("GET", url.String(), nil)
	//request.Header.Add("Proxy-Authorization", basic)
	
	dump, _ := httputil.DumpRequest(request, false)
	fmt.Println(string(dump))
	
	response, err := client.Do(request)
	checkError(err)
	
	fmt.Println("Read OK")
	
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}
	fmt.Println("Response OK")
	
	var buf [512]byte
	reader := response.Body
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Println(string(buf[0:n]))
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}