package main

import (
	"net"
	"strconv"
	"time"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func initMain() {
	// app started
	log.WithFields(log.Fields{"Booted": 1}).Info("Order")

	// start pubsub
	go initPubsub() 

	// start the web stuff
	go initHTTP()

	// connect to redis
	S := 0
	conn, err := net.Dial("udp", "redis:"+Conf.RedisPort)
	for retry := 1; err != nil; retry++ {
		S := strconv.Itoa(retry)
		log.WithFields(log.Fields{"Connected": 0, "Attempt": S}).Warn("Order:Redis")
		if retry > 3 {
			log.WithFields(log.Fields{"Connected": 0, "Attempt": 5}).Fatal("Order:Redis")
		}
		time.Sleep(5 * time.Second)
	}
	log.WithFields(log.Fields{"Connected": 1, "Attempt": S}).Info("Order:Redis")
	conn.Close()

	// start plugins
	go initPlugins()
}

func initPlugins() {

	if Conf.PluginDiscord == true {
		log.WithFields(log.Fields{"Enabled": 1}).Info("Order:Discord")
		go initDiscord()
	} else {
		log.WithFields(log.Fields{"Enabled": 0}).Info("Order:Discord")
	}

	if Conf.PluginHearbeat == true {
		log.WithFields(log.Fields{"Enabled": 1}).Info("Order:Heartbeat")
		go initHeartbeat()
	} else {
		log.WithFields(log.Fields{"Enabled": 0}).Info("Order:Heartbeat")
	}

	if Conf.PluginLogs == true {
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
