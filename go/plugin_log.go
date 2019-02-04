package main

import (
	"fmt"

	"github.com/hpcloud/tail"
	log "github.com/sirupsen/logrus"
)

func initLog() {
	// app started
	log.WithFields(log.Fields{"Started": 1}).Info("Order:Log")

	go watchNwnxeeLog()
}

func watchNwnxeeLog() {
	t, err := tail.TailFile("/logs/nwnx.txt", tail.Config{Follow: true})
	for line := range t.Lines {
		log.WithFields(log.Fields{"": 1}).Info("Order:Log:Nwnxee:" + line.Text)
	}

	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
