package tace

import (
	"context"
	"flag"
	"fmt"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	server = "server"
	client = "client"
)
var (
	actorKind  = flag.String("actor", "server", "server or client")
)

func main() {
	flag.Parse()

	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1.0, // sample all traces
		},
	}

	tracer, closer, err := cfg.New(*actorKind)
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	if *actorKind == server {
		runServer(tracer)
	} else {
		runClient(tracer)
	}

}

func getTime(w http.ResponseWriter, r *http.Request) {
	log.Print("Received getTime request")
	t := time.Now()
	ts := t.Format("Mon Jan _2 15:04:05 2006")
	io.WriteString(w, fmt.Sprintf("The time is %s", ts))
}

func runServer(tracer opentracing.Tracer) {
	http.HandleFunc("/gettime", getTime)
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", 8081),
		nethttp.Middleware(tracer, http.DefaultServeMux))
	if err != nil {
		log.Fatalf("Cannot start server: %s", err)
	}
}

func runClient(tracer opentracing.Tracer) {
	c := &http.Client{Transport: &nethttp.Transport{}}
	span := tracer.StartSpan(*actorKind)
	span.SetTag(string(ext.Component), client)
	span.SetTag("snlan-debug-id", "some-correlation-id")
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	req, err := http.NewRequest("GET", "http://localhost:8081/gettime", nil)
	if err != nil {
		onError(span, err)
		return
	}
	req = req.WithContext(ctx)
	req, ht := nethttp.TraceRequest(tracer, req)
	defer ht.Finish()

	res, err := c.Do(req)
	if err != nil {
		onError(span, err)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		onError(span, err)
		return
	}

	fmt.Printf("Received result: %s\n", string(body))
}

func onError(span opentracing.Span, err error) {
	// handle errors by recording them in the span
	span.SetTag(string(ext.Error), true)
	span.LogKV(otlog.Error(err))
	log.Print(err)
}
