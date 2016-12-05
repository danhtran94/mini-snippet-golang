package main

import (
	"net/rpc"
	"fmt"
	"log"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main()  {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	checkError("logging", err)
	
	args := Args{17, 8}
	
	var reply int
	err = client.Call("Operator.Multiply", args, &reply)
	checkError("arith error", err)
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)
	
	var reply2 int
	err = client.Call("Hello.Hi", args, &reply2)
	checkError("arith error", err)
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply2)
	
	var quot Quotient
	err = client.Call("Operator.Divide", args, &quot)
	checkError("arith error", err)
	fmt.Printf("Arith: %d / %d = %d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
	
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, ": ", err.Error())
	}
}
