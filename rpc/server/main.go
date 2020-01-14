package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type T struct {

}

func (t *T)SetPerson(person Person, p2 *Person) error {
	fmt.Printf("%#v\n", person)

	p2.Age=12
	p2.Name = person.Name
	return nil
}

type Person struct {
	Name string
	Age  int
}

func main() {
	rpcs := rpc.NewServer()
	err := rpcs.Register(new(T))
	if err != nil {
		panic(err)
	}
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		conn, err := l.Accept()
		if err == nil {
			go rpcs.ServeConn(conn)
		} else {
			break
		}
	}
	l.Close()
}

