package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func main() {
	zk, _, err := zk.Connect([]string{"www.snlan.top:2181"}, time.Second)
	if err != nil {
		panic(err)
	}

LOOP:
	data, _, events, err := zk.GetW("/dal_orm_release")
	fmt.Printf("data %s\n", string(data))
	ev := <- events
	fmt.Printf("ev %v\n", ev)
	goto LOOP

}
