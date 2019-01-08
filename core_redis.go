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
		"input",
		"debug",
		"github",
	)
	log.WithFields(log.Fields{"heartbeat": "1", "input": "1", "debug": "1", "github": "1"}).Info("Order:Redis:Pubsub:Subscribe")
	for {
		msg, _ := pubSub.ReceiveMessage()
		switch msg.Channel {
		case "heartbeat":

		case "input":

		case "debug":
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
