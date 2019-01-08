package main

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/hpcloud/tail"
	log "github.com/sirupsen/logrus"
)

func initLog() {
	// grab config
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	// app started
	log.WithFields(log.Fields{"Started": 1}).Info("Order:Log")

	go watchNwserverLog()
	go watchNwnxeeLog()
}

func watchNwserverLog() {
	t, err := tail.TailFile("/logs/nwserverLog1.txt", tail.Config{Follow: true})
	for line := range t.Lines {
		log.WithFields(log.Fields{"": 1}).Info("Order:Log:Nwserver:" + line.Text)
	}

	if err != nil {
		fmt.Printf("%+v\n", err)
	}
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
