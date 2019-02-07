package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func initMain() {
	// app started
	log.WithFields(log.Fields{"Booted": 1}).Info("Order")
	go InitAPI()
	go initPlugins()
}

func initPlugins() {
	if os.Getenv("NWN_ORDER_DISCORD_ENABLED") == "1" {
		log.WithFields(log.Fields{"Enabled": 1}).Info("Order:Discord")
		go initDiscord()
	} else {
		log.WithFields(log.Fields{"Enabled": 0}).Info("Order:Discord")
	}

	if os.Getenv("NWN_ORDER_HEARTBEAT_ENABLED") == "1" {
		log.WithFields(log.Fields{"Enabled": 1}).Info("Order:Heartbeat")
		go initHeartbeat()
	} else {
		log.WithFields(log.Fields{"Enabled": 0}).Info("Order:Heartbeat")
	}

	if os.Getenv("NWN_ORDER_LOG_ENABLED") == "1" {
		log.WithFields(log.Fields{"Enabled": 1}).Info("Order:Logs")
		go initLog()
	} else {
		log.WithFields(log.Fields{"Enabled": 0}).Info("Order:Logs")
	}
}

func checkENV() (err error) {
	// required env
	if os.Getenv("NWN_ORDER_PORT") == "" {
		log.WithFields(log.Fields{"NWN_ORDER_PORT": "nil"}).Fatal("Order:Core:Env")
	}
	if os.Getenv("NWN_ORDER_MODULE_NAME") == "" {
		log.WithFields(log.Fields{"NWN_ORDER_MODULE_NAME": "nil"}).Fatal("Order:Core:Env")
	}
	if os.Getenv("NWN_ORDER_REDIS_HOST") == "" {
		log.WithFields(log.Fields{"NWN_ORDER_REDIS_HOST": "nil"}).Fatal("Order:Core:Env")
	}
	if os.Getenv("NWN_ORDER_REDIS_PORT") == "" {
		log.WithFields(log.Fields{"NWN_ORDER_REDIS_PORT": "nil"}).Fatal("Order:Core:Env")
	}
	return
}

func main() {
	// all the cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// load the core config here
	err := checkENV()
	if err != nil {
		return
	}

	// Initialize order
	go initMain()

	// forever
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-done
}
