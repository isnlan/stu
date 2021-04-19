package main

import (
	"context"
	"fmt"
	"log"
	"stu/grpc-test/helloworld"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		c := helloworld.NewGreeterClient(conn)
		go func(id int) {
			for j := 0; j < 1000; j++ {
				c.SayHello(context.Background(), &helloworld.HelloRequest{
					Name: fmt.Sprintf("id%d", id),
				})
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	time.Sleep(time.Minute * 10)
}
