package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func proxy(c *gin.Context) {
	remote, err := url.Parse("http:/127.0.0.1:8500")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	//Define the director func
	//This is a good place to log, for example
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {
	router := gin.Default()

	//Create a catchall route
	router.Any("/*proxyPath", proxy)

	//router.Any("/consul", func(c *gin.Context) {
	//	//action := c.Param("action")
	//	//fmt.Println("...", action)
	//	//action = strings.ReplaceAll(action, "/proxy", "")
	//	url, _ := url.Parse("127.0.0.1:8500")
	//	proxy := httputil.NewSingleHostReverseProxy(url)
	//	//director := func(req *http.Request) {
	//	//	req.URL.Scheme = "http"
	//	//	req.URL.Host = "127.0.0.1:8500"
	//	//	req.URL.Path, req.URL.RawPath = "/ui", "/ui"
	//	//	if _, ok := req.Header["User-Agent"]; !ok {
	//	//		// explicitly disable User-Agent so it's not set to default value
	//	//		req.Header.Set("User-Agent", "")
	//	//	}
	//	//}
	//
	//	//proxy := &httputil.ReverseProxy{Director: director}
	//	proxy.ServeHTTP(c.Writer, c.Request)
	//})
	//
	//router.Any("/ui", func(c *gin.Context) {
	//	//action := c.Param("action")
	//	//fmt.Println("...", action)
	//	//action = strings.ReplaceAll(action, "/proxy", "")
	//	url, _ := url.Parse("127.0.0.1:8500")
	//	proxy := httputil.NewSingleHostReverseProxy(url)
	//	//director := func(req *http.Request) {
	//	//	req.URL.Scheme = "http"
	//	//	req.URL.Host = "127.0.0.1:8500"
	//	//	req.URL.Path, req.URL.RawPath = "/ui", "/ui"
	//	//	if _, ok := req.Header["User-Agent"]; !ok {
	//	//		// explicitly disable User-Agent so it's not set to default value
	//	//		req.Header.Set("User-Agent", "")
	//	//	}
	//	//}
	//
	//	//proxy := &httputil.ReverseProxy{Director: director}
	//	proxy.ServeHTTP(c.Writer, c.Request)
	//})

	//router.GET("/ui/*action", func(c *gin.Context) {
	//	action := c.Param("action")
	//	fmt.Println("...", action)
	//	//action = strings.ReplaceAll(action, "/proxy", "")
	//	director := func(req *http.Request) {
	//		req.URL.Scheme = "http"
	//		req.URL.Host = "127.0.0.1:8500"
	//		req.URL.Path, req.URL.RawPath = action, action
	//		if _, ok := req.Header["User-Agent"]; !ok {
	//			// explicitly disable User-Agent so it's not set to default value
	//			req.Header.Set("User-Agent", "")
	//		}
	//	}
	//
	//	proxy := &httputil.ReverseProxy{Director: director}
	//	proxy.ServeHTTP(c.Writer, c.Request)
	//})

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

	router.Run(":8080")
}

func HandleNotFound(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "404 page not found",
		"app":     "blockchain-api",
		"request": c.Request.Method + " " + c.Request.URL.String(),
	})
}
