package main

import (
	"github.com/golang/glog"
)

func main() {
	glog.V(1).Info("Completed successfully.")
	glog.Flush()
}
