package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func initMain() {
	// app started
	log.WithFields(log.Fields{"Booted": 1}).Info("Order")

	go initConf()
	log.WithFields(log.Fields{"Config": 1}).Info("Order")

	// start the web stuff
	go initHTTP()
	log.WithFields(log.Fields{"API": 1}).Info("Order")

	// start pubsub
	go initPubsub() 
	log.WithFields(log.Fields{"Pubsub": 1}).Info("Order")

	// connect to redis

	// start plugins
	//go initPlugins()
}

func initPlugins() {
	c := Config{}
	if c.PluginDiscord == true {
		log.WithFields(log.Fields{"Enabled": 1}).Info("Order:Discord")
		go initDiscord()
	} else {
		log.WithFields(log.Fields{"Enabled": 0}).Info("Order:Discord")
	}

	if c.PluginHearbeat == true {
		log.WithFields(log.Fields{"Enabled": 1}).Info("Order:Heartbeat")
		go initHeartbeat()
	} else {
		log.WithFields(log.Fields{"Enabled": 0}).Info("Order:Heartbeat")
	}

	if c.PluginLogs == true {
		log.WithFields(log.Fields{"Enabled": 1}).Info("Order:Logs")
		go initLog()
	} else {
		log.WithFields(log.Fields{"Enabled": 0}).Info("Order:Logs")
	}
}

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	go initMain()
	<-done // Block forever, not sure if this is best practice.
}
