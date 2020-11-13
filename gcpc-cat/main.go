package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"stu/gcpc-cat/protos"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8081"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protos.NewBehaviorLogClient(conn)

	// Contact the server and print out its response.
	var list []*protos.Resource
	list = append(list, &protos.Resource{
		Id:   "myid",
		Type: "DataService.ID",
	})

	for i := 0; i < 2; i++ {
		l := protos.Record{
			Service:   "stu client",
			NetworkId: "fabric",
			Channel:   "service channel",
			Contract:  "kyc",
			Proxy:     "kd-proxy",
			Bduid:     "user1-bduid",
			Operation: "DataService.Create",
			Resources: list,
			TxId:      fmt.Sprintf("tx-a_%d", i),
			SpanId:    "fdsfafd",
			TraceId:   "fdsafdaf",
		}

		bytes, _ := json.Marshal(&l)
		fmt.Println(string(bytes))

		_, err = c.Recode(context.Background(), &l)
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		time.Sleep(time.Second)
	}

}
