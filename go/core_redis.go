package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// IRedisClient interface
type IRedisClient interface {
	HMGet(key string, fields ...string) (string, error)
	HMSet(key string, fields map[string]interface{}) (string, error)
	HDel(key string, fields ...string) (int64, error)
	HIncrBy(key, field string, incr int64) (int64, error)
	HGet(key, field string) (string, error)
	HSet(key, field string, value interface{}) (bool, error)
}

// RedisClient struct
type RedisClient struct {
	client *redis.Client
}

// NewRedisClient func
func NewRedisClient(address string, password string) IRedisClient {
	return &RedisClient{client: redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       0,        // use default DB
	})}
}

// HMGet func
func (c *RedisClient) HMGet(key string, fields ...string) (string, error) {
	var result string
	response, err := c.client.HMGet(key, fields...).Result()
	if err == nil && response != nil && len(response) > 0 && response[0] != nil {
		result = response[0].(string)
	}

	return result, err
}

// HMSet func
func (c *RedisClient) HMSet(key string, fields map[string]interface{}) (string, error) {
	var result string

	response, err := c.client.HMSet(key, fields).Result()

	if err == nil {
		result = response
	}

	return result, err
}

// HDel func
func (c *RedisClient) HDel(key string, fields ...string) (int64, error) {
	var result int64

	response, err := c.client.HDel(key, fields...).Result()

	if err == nil {
		result = response
	}

	return result, err
}

// HIncrBy func
func (c *RedisClient) HIncrBy(key, field string, incr int64) (int64, error) {
	var result int64

	response, err := c.client.HIncrBy(key, field, incr).Result()

	if err == nil {
		result = response
	}

	return result, err
}

// HGet func
func (c *RedisClient) HGet(key string, field string) (string, error) {
	var result string
	result, err := c.client.HGet(key, field).Result()

	if err == redis.Nil {
		return result, nil
	}

	return result, err
}

// HSet func
func (c *RedisClient) HSet(key string, field string, value interface{}) (bool, error) {
	var result bool

	result, err := c.client.HSet(key, field, value).Result()

	return result, err
}

func initPubsub() {
	addr := os.Getenv("NWN_ORDER_REDIS_HOST") + ":" + os.Getenv("NWN_ORDER_REDIS_PORT")
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
		log.WithFields(log.Fields{"Connected": "0", "Please confirm redis is connected": "1"}).Fatal("Order:Redis")
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
	r := redis.NewClient(&redis.Options{Addr: os.Getenv("NWN_ORDER_REDIS_HOST") + ":" + os.Getenv("NWN_ORDER_REDIS_PORT")})
	err := r.Publish(PubsubChannel, PubsubMessage).Err()
	if err != nil {
		log.WithFields(log.Fields{"Channel": PubsubChannel, "Message": PubsubMessage}).Error(LogMessage)
		panic(err)
	}
	log.WithFields(log.Fields{"Channel": PubsubChannel, "Message": PubsubMessage}).Info(LogMessage)
}
