package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello Skaffold cicd!")
		time.Sleep(time.Second * 2)
	}
}
