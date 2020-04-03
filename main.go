package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)
type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Arith int

type ArithAddResp struct {
	Id     interface{} `json:"id"`
	Result Reply       `json:"result"`
	Error  interface{} `json:"error"`
}

func (t *Arith) Add(args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

func (t *Arith) Mul(args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func (t *Arith) Div(args *Args, reply *Reply) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	reply.C = args.A / args.B
	return nil
}


func main() {
	cli, srv := net.Pipe()
	defer cli.Close()
	go jsonrpc.ServeConn(srv)
	dec := json.NewDecoder(cli)

	// Send hand-coded requests to server, parse responses.
	for i := 0; i < 10; i++ {
		fmt.Fprintf(cli, `{"method": "Arith.Add", "id": "\u%04d", "params": [{"A": %d, "B": %d}]}`, i, i, i+1)
		var resp ArithAddResp
		err := dec.Decode(&resp)
		if err != nil {
			panic(err)
		}
		if resp.Error != nil {
			panic(resp)
		}
		if resp.Id.(string) != string(i) {
			panic(resp)
		}
		if resp.Result.C != 2*i+1 {
			panic(resp)
		}
	}
}

