package main

import (
	"fmt"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	name string
}

func main() {
	s := "Md5"
	lower := strings.ToLower(s)
	fmt.Println(lower, 1)
}
