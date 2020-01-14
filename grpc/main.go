package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"stu/grpc/kvdb"
)

type KVDBService struct {
}

func (s *KVDBService) Set(context.Context, *kvdb.SetKV) (*kvdb.Empty, error) {
	//log.Println("-")
	return &kvdb.Empty{}, nil
}

const (
	port = ":9090"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	kvdb.RegisterKvDBServer(s, &KVDBService{})
	for k, v := range s.GetServiceInfo() {
		fmt.Println(k, v)

	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
