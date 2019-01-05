package main

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/go-redis/redis"
)

var (
	//RedisClient is a var
	RedisClient *redis.Client
)

func startPubsub() {
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

	if cfg.HbVerbose == true {
		fmt.Println(LogMessage)
	}
}
