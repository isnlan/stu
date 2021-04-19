package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const sockAddr = "/data/gopath/src/stu/unixsocket/http.sock"

func main() {
	router := gin.New()
	router.GET("/testGet", handlerGet)
	os.Remove(sockAddr)
	unixAddr, err := net.ResolveUnixAddr("unix", sockAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	listener, err := net.ListenUnix("unix", unixAddr)
	if err != nil {
		fmt.Println("listening error:", err)
	}
	fmt.Println("listening...")
	http.Serve(listener, router)
}

func handlerGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"resp": "ok",
	})
}
