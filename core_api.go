package main

import "github.com/go-redis/redis"

// ModuleBootTime func
func ModuleBootTime() string {
	r := redis.NewClient(&redis.Options{Addr: "redis:" + Conf.RedisPort})
	rkey := Conf.ModuleName + ":server"
	value, _ := r.HGet(rkey, "BootTime").Result()
	return value
}

// ModuleBootDate func
func ModuleBootDate() string {
	r := redis.NewClient(&redis.Options{Addr: "redis:" + Conf.RedisPort})
	rkey := Conf.ModuleName + ":server"
	value, _ := r.HGet(rkey, "BootDate").Result()
	return value
}

// ModulePlayers func
func ModulePlayers() string {
	r := redis.NewClient(&redis.Options{Addr: "redis:" + Conf.RedisPort})
	rkey := Conf.ModuleName + ":server"
	value, _ := r.HGet(rkey, "Online").Result()
	return value
}
