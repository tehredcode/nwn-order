package main

import (
	"github.com/go-redis/redis"
)

// ModuleBootTime func
func ModuleBootTime() string {
	c := Config{}
	r := redis.NewClient(&redis.Options{Addr: "redis:" + c.RedisPort})
	rkey := c.ModuleName + ":server"
	value, _ := r.HGet(rkey, "BootTime").Result()
	return value
}

// ModuleBootDate func
func ModuleBootDate() string {
	c := Config{}
	r := redis.NewClient(&redis.Options{Addr: "redis:" + c.RedisPort})
	rkey := c.ModuleName + ":server"
	value, _ := r.HGet(rkey, "BootDate").Result()
	return value
}

// ModulePlayers func
func ModulePlayers() string {
	c := Config{}
	r := redis.NewClient(&redis.Options{Addr: "redis:" + c.RedisPort})
	rkey := c.ModuleName + ":server"
	value, _ := r.HGet(rkey, "Online").Result()
	return value
}
