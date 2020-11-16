package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/ide/*action", func(c *gin.Context) {
		action := c.Param("action")
		fmt.Println("...", action)
		action = strings.ReplaceAll(action, "/proxy", "")
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = "dapp-5face284f6add4bd611d2264.dapp-project:3000"
			req.URL.Path, req.URL.RawPath = action, action
			if _, ok := req.Header["User-Agent"]; !ok {
				// explicitly disable User-Agent so it's not set to default value
				req.Header.Set("User-Agent", "")
			}
		}

		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	//router.GET("/ide", func(c *gin.Context) {
	//	director := func(req *http.Request) {
	//		req.URL.Scheme = "http"
	//		req.URL.Host = "dapp-5face284f6add4bd611d2264.dapp-project:3000"
	//		req.URL.Path, req.URL.RawPath = "", ""
	//		//if targetQuery == "" || req.URL.RawQuery == "" {
	//		//	req.URL.RawQuery = targetQuery + req.URL.RawQuery
	//		//} else {
	//		//	req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
	//		//}
	//		if _, ok := req.Header["User-Agent"]; !ok {
	//			// explicitly disable User-Agent so it's not set to default value
	//			req.Header.Set("User-Agent", "")
	//		}
	//	}
	//	proxy := &httputil.ReverseProxy{Director: director}
	//	proxy.ServeHTTP(c.Writer, c.Request)
	//})

	router.Run()
}

func HandleNotFound(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "404 page not found",
		"app":     "blockchain-api",
		"request": c.Request.Method + " " + c.Request.URL.String(),
	})
}
