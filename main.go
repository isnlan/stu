package main

import (
	"log"
	"net/http"
	"net/url"
	"stu/wsproxy"
)

func main() {
	u, err := url.Parse("ws://dapp-5face284f6add4bd611d2264.dapp-project:3000/services")
	if err != nil {
		log.Fatalln(err)
	}

	err = http.ListenAndServe(":3001", wsproxy.NewProxy(u))
	if err != nil {
		log.Fatalln(err)
	}
}
