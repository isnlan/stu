package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("bazel build golang")
		time.Sleep(time.Second * 2)
	}
}
