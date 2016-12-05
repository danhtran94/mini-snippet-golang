package main

import (
	"fmt"
	"net/rpc"
	"errors"
	"net/http"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith interface{
	Multiply(args *Args, reply *int) error
	Divide(args *Args, quo *Quotient) error
}

type Operator struct {}
type Hello struct {}

func (Operator) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (Operator) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("devide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func (Hello) Hi(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func main() {
	var rpcHandler Arith = Operator{}
	var rpcHello2 Hello = Hello{}
	rpc.Register(rpcHandler)
	rpc.Register(rpcHello2)
	rpc.HandleHTTP()
	
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

