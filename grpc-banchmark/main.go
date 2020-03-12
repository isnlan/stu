package main

import (
	"fmt"
	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
	"os"
	"stu/grpc-banchmark/helloworld"
)

func main()  {
	report, err := runner.Run(
		"helloworld.Greeter.SayHello",
		"localhost:50051",
		runner.WithProtoFile("/Users/snlan/go/src/stu/grpc-banchmark/helloworld/helloworld.proto", []string{}),
		runner.WithData(helloworld.HelloRequest{
			Name:                 "job",
		}),
		runner.WithInsecure(true),
		runner.WithConcurrency(10000),
		runner.WithConnections(10),
	)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	printer := printer.ReportPrinter{
		Out:    os.Stdout,
		Report: report,
	}

	printer.Print("pretty")
}
