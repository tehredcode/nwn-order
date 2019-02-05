package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/codegangsta/martini"
	"github.com/garyburd/redigo/redis"
	"github.com/martini-contrib/render"
	log "github.com/sirupsen/logrus"
)

var (
	redisAddress   = flag.String("redis-address", ":6379", "Address to the Redis server")
	maxConnections = flag.Int("max-connections", 10, "Max connections to Redis")
)

func InitAPI() {
	martini.Env = martini.Prod

	flag.Parse()

	redisPool := redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", *redisAddress)

		if err != nil {
			return nil, err
		}

		return c, err
	}, *maxConnections)

	defer redisPool.Close()

	m := martini.Classic()

	m.Map(redisPool)

	m.Use(render.Renderer())

	m.Get("/stats", func(r render.Render, pool *redis.Pool, params martini.Params) {
		c := pool.Get()
		defer c.Close()

		k := "Order:server"
		value1 := "Order"
		value2, err := redis.String(c.Do("HGET", k, "BootTime"))
		value3, err := redis.String(c.Do("HGET", k, "BootDate"))

		if err != nil {
			message := fmt.Sprintf("Could not GET %s", "key")

			r.JSON(400, map[string]interface{}{
				"status":  "ERR",
				"message": message})
		} else {
			r.JSON(200, map[string]interface{}{
				"status":     "OK",
				"ModuleName": value1,
				"BootTime":   value2,
				"BootDate":   value3,
			})
		}
	})

	m.Run()
}

func initMain() {
	// app started
	log.WithFields(log.Fields{"Booted": 1}).Info("Order")
	go InitAPI()
}

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	go initMain()
	<-done
}
