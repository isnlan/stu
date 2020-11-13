package main

import (
	"context"
	"fmt"
	"log"
	"stu/grpc-test/helloworld"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	go func() {
		time.Sleep(time.Second * 2)
		conn.Close()
	}()

	reply, err := c.SayHello(context.Background(), &helloworld.HelloRequest{
		Name: "snlna",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(reply)

}
