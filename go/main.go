package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/urothis/nwn-order/go/log"

	log "github.com/sirupsen/logrus"
)

func initMain() {
	// app started
	log.WithFields(log.Fields{"Booted": 1}).Info("Order")

	// grab redis client
	client := InitRedisClient()

	// start the web stuff
	go initHTTP()
	log.WithFields(log.Fields{"API": 1}).Info("Order")

	// start pubsub
	go initPubsub()
	log.WithFields(log.Fields{"Pubsub": 1}).Info("Order")

	// start plugins
	go initPlugins()
}

func initPlugins() {
	if os.Getenv("NWN_ORDER_PLUGIN_DISCORD_ENABLED") == "1" {
		log.WithFields(log.Fields{"Enabled": 1}).Info("Order:Discord")
		go initDiscord()
	} else {
		log.WithFields(log.Fields{"Enabled": 0}).Info("Order:Discord")
	}

	if os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_ENABLED") == "1" {
		log.WithFields(log.Fields{"Enabled": 1}).Info("Order:Heartbeat")
		go initHeartbeat()
	} else {
		log.WithFields(log.Fields{"Enabled": 0}).Info("Order:Heartbeat")
	}

	if os.Getenv("NWN_ORDER_PLUGIN_LOG_ENABLED") == "1" {
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
	<-done
}
