package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fileServer := http.FileServer(http.Dir("/var/www/"))
	http.Handle("/", fileServer)
	
	http.HandleFunc("/cgi-bin/printenv", printEnv)
	err := http.ListenAndServe(":8000", nil)
	checkError(err)
}

func printEnv(rw http.ResponseWriter, req *http.Request) {
	env := os.Environ()
	rw.Write([]byte("<h1>Environment</h1>"))
	for _, v := range env {
		rw.Write([]byte("<p>"+ v +"</p>"));
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}

