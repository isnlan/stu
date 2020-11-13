package main

import (
	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	hook, err := logrustash.NewHook("tcp", "127.0.0.1:5000", "stu")
	if err != nil {
		log.Fatal(err)
	}

	log.Hooks.Add(hook)
	ctx := log.WithFields(logrus.Fields{
		"method": "mainfsf",
	})
	ctx.Info("Hello World!")
	ctx.Warn("wind")
}
