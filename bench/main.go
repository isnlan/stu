package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"
)

type CommandSet struct {
	Set struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
}

type CommandRemove struct {
	Remove struct {
		Key string `json:"key"`
	}
}

type CommandGet struct {
	Get struct {
		Key string `json:"key"`
	}
}

func main() {
	workers := 100
	connChannel := make(chan net.Conn, workers)
	start := time.Now()
	for i := 0; i < workers; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:4000")
		if err != nil {
			panic(err)
		}
		connChannel <- conn
	}
	close(connChannel)

	var wg = &sync.WaitGroup{}
	for conn := range connChannel {
		wg.Add(1)
		go func(c net.Conn) {
			doWork(c)
			wg.Done()
		}(conn)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}

func doWork(conn net.Conn) {
	var cmdList [][]byte
	for i := 0; i < 10000; i++ {
		cs := CommandSet{Set: struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}{Key: fmt.Sprintf("key-%s-%d", conn.LocalAddr().String(), i), Value: fmt.Sprintf("value-%d", i)}}

		bytes, err := json.Marshal(cs)
		check(err)
		cmdList = append(cmdList, bytes)
	}

	for i := 0; i < 10000; i++ {
		cs := CommandGet{Get: struct {
			Key string `json:"key"`
		}{Key: fmt.Sprintf("key-%s-%d", conn.LocalAddr().String(), i)}}

		bytes, err := json.Marshal(cs)
		check(err)
		cmdList = append(cmdList, bytes)
	}

	for i := 0; i < 10000; i++ {
		cs := CommandRemove{Remove: struct {
			Key string `json:"key"`
		}{Key: fmt.Sprintf("key-%s-%d", conn.LocalAddr().String(), i)}}

		bytes, err := json.Marshal(cs)
		check(err)
		cmdList = append(cmdList, bytes)
	}

	for _, cmd := range cmdList {
		_, err := conn.Write(cmd)
		if err != nil {
			panic(err)
		}
	}

}

func check(err error) {
	if err != nil {
		panic(err)
	}

}
