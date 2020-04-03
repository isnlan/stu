package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"time"
)

type Calc struct{}

type Args struct {
	A  float64 `json:"a"`
	B  float64 `json:"b"`
	Op string  `json:"op"`
}

type Reply struct {
	Msg  string  `json:"msg"`
	Data float64 `json:"data"`
}

func Call(args Args, reply *Reply) error {
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalln("dailing error: ", err)
		return err
	}

	defer conn.Close()

	// 调用远程的Calc的Compute方法
	err = conn.Call("Calc.Compute", args, &reply)
	return err
}

func main()  {
		time.Sleep(time.Second*2)
		args := Args{
			A:  1,
			B:  2,
			Op: "+",
		}

		var reply Reply
		err := Call(args, &reply)
		if err != nil {
			panic(err)
		}
		fmt.Println(reply)
}

