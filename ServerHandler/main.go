package main

import (
	"net/http"
)

func main() {
	myHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusNoContent)
	})
	
	http.ListenAndServe(":2000", myHandler)
}

