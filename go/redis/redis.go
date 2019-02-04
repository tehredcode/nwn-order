package redis

import (
	"os"

	"github.com/go-redis/redis"
)

// IRedisClient interface
type IClient interface {
	HMGet(key string, fields ...string) (string, error)
	HMSet(key string, fields map[string]interface{}) (string, error)
	HDel(key string, fields ...string) (int64, error)
	HIncrBy(key, field string, incr int64) (int64, error)
	HGet(key, field string) (string, error)
	HSet(key, field string, value interface{}) (bool, error)
}

// RedisClient struct
type Client struct {
	client *redis.Client
}

// InitRedisClient func
func InitClient() IRedisClient {
	return &RedisClient{client: redis.NewClient(&redis.Options{
		Addr:     os.Getenv("NWN_ORDER_REDIS_HOST"),
		Password: os.Getenv("NWN_ORDER_REDIS_PW"),
		DB:       0, // use default DB
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
