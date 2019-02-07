package main

import (
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
)

var (
	// RedisPool is the active redis connection
	RedisPool *redis.Pool
)

func hgetRediskeyString(key string, field string) (string, error) {
	redisCon := RedisPool.Get()
	defer redisCon.Close()
	result, err := redis.String(redisCon.Do("hget", key, field))
	if err != nil {
		log.Warnln("failed to perform hget operation on redis", err)
	}
	return result, err
}

func hgetRediskeyInt(key string, field string) (int, error) {
	redisCon := RedisPool.Get()
	defer redisCon.Close()
	result, err := redis.Int(redisCon.Do("hget", key, field))
	if err != nil {
		log.Warnln("failed to perform hget operation on redis", err)
	}
	return result, err
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     300,
		MaxActive:   600,
		IdleTimeout: 20000 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func sendPubsub(LogMessage string, PubsubChannel string, PubsubMessage string) {
	r := RedisPool.Get()
	err := r.Send(PubsubChannel, PubsubMessage)
	if err != nil {
		log.WithFields(log.Fields{"Channel": PubsubChannel, "Message": PubsubMessage}).Fatal(LogMessage)
	}

	if os.Getenv("NWN_ORDER_REDIS_PUBSUB_VERBOSE") == "1" {
		log.WithFields(log.Fields{"Channel": PubsubChannel, "Message": PubsubMessage}).Info(LogMessage)
	}
}
