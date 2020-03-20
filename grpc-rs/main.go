package main

import (
	"context"
	"fmt"
	proto2 "github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	proto "stu/grpc-rs/protos"
)


func runRegister(client proto.ComponentClient) {
	register, err := client.Register(context.Background())
	if err != nil {
		panic(err)
	}
	go func() {
		for i:=0 ; i < 3; i++ {
			reg := &proto.ComponentContractRegister{
				Name:                 "kv",
				Decorations:          nil,
			}
			bytes, err := proto2.Marshal(reg)
			if err != nil {
				panic(err)
			}
			err = register.Send(&proto.Message{
				MessageType:          proto.Message_COMPONENT_CONTRACT_REGISTER,
				CorrelationId:        "",
				Content:              bytes,
			})
			if err != nil {
				panic(err)
			}
		}
	}()

	for {
		recv, err := register.Recv()
		if err == io.EOF {
			fmt.Println("close")
			return
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("recv message :%v\n", recv)
	}
}


//go:generate protoc -I protos/ protos/block.proto protos/block.proto protos/common.proto protos/consensus.proto protos/message.proto protos/peer.proto --go_out=plugins=grpc:protos
func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewComponentClient(conn)

	mainPing()
	// RecordRoute
	runRegister(client)
}

func mainPing() {
	conn, err := grpc.Dial("127.0.0.1:8090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewEndorserClient(conn)
	ping, err := client.Ping(context.Background(), &proto.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("--> %+v\n", ping)

}
