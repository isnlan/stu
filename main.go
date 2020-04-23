package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"reflect"
	"stu/grpc-test/helloworld"
)

func print(msg proto.Message)  {
	fmt.Println(reflect.TypeOf(msg).String())
}

func main() {
	req := helloworld.HelloRequest{Name: "s"}
	print(&req)
}

