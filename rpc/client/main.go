package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func connect() *rpc.Client {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

func main() {
	client := connect()

	p := Person{Name:"ll", Age:340}
	var reply Person
	err := client.Call("T.SetPerson", &p, &reply)
	if err != nil {
		panic(err)
	}


	fmt.Println(reply)
	time.Sleep(time.Minute*10)
}

func main1() {
	var opts []grpc.DialOption

	f := "/opt/gopath/src/github.com/hyperledger/fabric/examples/e2e_cli/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.crt"
	c, err := credentials.NewClientTLSFromFile(f, "peer0.org1.example.com")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	}
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithTransportCredentials(c))
	ctx, canel := context.WithTimeout(context.Background(), time.Second*3)
	defer canel()
	conn, err := grpc.DialContext(ctx, "127.0.0.1:7051", opts...)
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
		return
	}

	println("ok", conn)

}

