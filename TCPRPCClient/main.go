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
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := Args{17, 8}
	reply := new(int)
	err = client.Call("Arith.Multiply", args, reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, *reply)
	
	quot := new(Quotient)
	err = client.Call("Arith.Divide", args, quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d / %d = %d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}

