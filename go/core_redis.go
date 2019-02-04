package main

import (
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"github.com/caarlos0/env"
)

// RedisClient struct
type RedisClient struct {
	Client *redis.Client
}

// NewRedisClient test
func NewRedisClient() (*RedisClient, error) {
	c := Config{}
	err := env.Parse(&c)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", c)

	conn := redis.NewClient(&redis.Options{
		Addr: c.RedisHost+":"+c.RedisPort,

	})
	_, err = conn.Ping().Result()
	if err != nil {
		log.WithFields(log.Fields{"Ping": 1, "pong": 1}).Info("Order:Redis")
		return nil, err
	}

	r := RedisClient{
		Client: conn,
	}

	defer conn.Close()
	return &r, nil
}

// HGet is a thin wrapper around redis.HGet of `github.com/go-redis/redis`.
func (r *RedisClient) HGet(key, field string) (string, error) {
	value, err := r.Client.HGet(key, field).Result()

	return value, err
}

// HSet is a thin wrapper around redis.HSet of `github.com/go-redis/redis`.
func (r *RedisClient) HSet(key, field, value string) (bool, error) {
	result, err := r.Client.HSet(key, field, value).Result()

	return result, err
}

// HExists is a thin wrapper around redis.HExists of `github.com/go-redis/redis`.
func (r *RedisClient) HExists(key, field string) (bool, error) {
	result, err := r.Client.HExists(key, field).Result()

	return result, err
}

func initPubsub() {
	c := Config{}
	err := env.Parse(&c)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	addr := c.RedisHost+":"+c.RedisPort
	r := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	
	pong, err := r.Ping().Result()
	fmt.Println(pong, err)

	defer r.Close()

	pubSub := r.Subscribe()
	err = pubSub.Subscribe(
		"Discord:Out",
		"Log:Debug",
		"Log:Info",
		"Log:Warning",
		"Log:Fatal",
	)
	if err != nil {
    	log.WithFields(log.Fields{"Connected": "0", "Please confirm redis is connected":"1"}).Fatal("Order:Redis")
	}
	defer pubSub.Close()
	log.WithFields(log.Fields{"Discord:Out": "1", "Log:Debug": "1", "Log:Info": "1", "Log:Warning": "1", "Log:Fatal": "1"}).Info("Order:Redis:Pubsub:Subscribe")
	
	for {
		msg, err := pubSub.ReceiveMessage()
		if err != nil {
			log.WithFields(log.Fields{"Connected": "0"}).Warn("Order:Redis:Pubsub")
		}
		switch msg.Channel {
		case "Discord:Out":
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
		fmt.Println(msg.Channel, msg.Payload)
	}
}

func sendPubsub(LogMessage string, PubsubChannel string, PubsubMessage string) {
	c := Config{}
	r := redis.NewClient(&redis.Options{Addr: "redis:" + c.ModuleName})
	err := r.Publish(PubsubChannel, PubsubMessage).Err()
	if err != nil {
		log.WithFields(log.Fields{"Channel": PubsubChannel, "Message": PubsubMessage}).Error(LogMessage)
		panic(err)
	}
	log.WithFields(log.Fields{"Channel": PubsubChannel, "Message": PubsubMessage}).Info(LogMessage)
}
