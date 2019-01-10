package main

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

func initPubsub() {
	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + cfg.RedisPort,
	})
	defer client.Close()

	pubSub := client.Subscribe(
		"heartbeat",
		"github",
		"Log:Debug",
		"Log:Info",
		"Log:Warning",
		"Log:Fatal",
	)
	log.WithFields(log.Fields{"heartbeat": "1", "github": "1", "Log:Debug": "1", "Log:Info": "1", "Log:Warning": "1", "Log:Fatal": "1"}).Info("Order:Redis:Pubsub:Subscribe")
	for {
		msg, _ := pubSub.ReceiveMessage()
		switch msg.Channel {
		case "heartbeat":
			log.WithFields(log.Fields{"Pubsub": "1", "Channel": msg.Channel, "Message": msg.Payload}).Info("Order:Pubsub")
		case "Log:Debug":
			log.WithFields(log.Fields{"Pubsub": "1", "Channel": msg.Channel, "Message": msg.Payload}).Info("Order:Pubsub")
		case "Log:Info":
			log.WithFields(log.Fields{"Pubsub": "1", "Channel": msg.Channel, "Message": msg.Payload}).Info("Order:Pubsub")
		case "Log:Warning":
			log.WithFields(log.Fields{"Pubsub": "1", "Channel": msg.Channel, "Message": msg.Payload}).Info("Order:Pubsub")
		case "Log:Fatal":
			log.WithFields(log.Fields{"Pubsub": "1", "Channel": msg.Channel, "Message": msg.Payload}).Info("Order:Pubsub")
		}
	}
}

func sendPubsub(LogMessage string, PubsubChannel string, PubsubMessage string) {
	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + cfg.RedisPort,
	})
	defer client.Close()

	if err := client.Publish(PubsubChannel, PubsubMessage).Err(); err != nil {
		panic(err)
	}

	log.WithFields(log.Fields{"Channel": PubsubChannel, "Message": PubsubMessage}).Info(LogMessage)
}
