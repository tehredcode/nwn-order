package main

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// RedisClient struct
type RedisClient struct {
	Client *redis.Client
}

// NewRedisClient test
func NewRedisClient() (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + Conf.ModuleName,
	})

	_, err := client.Ping().Result()

	if err != nil {
		return nil, err
	}

	r := RedisClient{
		Client: client,
	}

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
	r := redis.NewClient(&redis.Options{Addr: "redis:" + Conf.ModuleName})
	defer r.Close()

	pubSub := r.Subscribe(
		"Discord:Out",
		"Log:Debug",
		"Log:Info",
		"Log:Warning",
		"Log:Fatal",
	)
	log.WithFields(log.Fields{"Discord:Out": "1", "Log:Debug": "1", "Log:Info": "1", "Log:Warning": "1", "Log:Fatal": "1"}).Info("Order:Redis:Pubsub:Subscribe")
	for {
		msg, _ := pubSub.ReceiveMessage()
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
	}
}

func sendPubsub(LogMessage string, PubsubChannel string, PubsubMessage string) {
	r := redis.NewClient(&redis.Options{Addr: "redis:" + Conf.ModuleName})
	err := r.Publish(PubsubChannel, PubsubMessage).Err()
	if err != nil {
		log.WithFields(log.Fields{"Channel": PubsubChannel, "Message": PubsubMessage}).Error(LogMessage)
		panic(err)
	}
	log.WithFields(log.Fields{"Channel": PubsubChannel, "Message": PubsubMessage}).Info(LogMessage)
}
