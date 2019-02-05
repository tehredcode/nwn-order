package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/codegangsta/martini"
	"github.com/garyburd/redigo/redis"
	"github.com/martini-contrib/render"
	log "github.com/sirupsen/logrus"
)

var (
	redisPool *redis.Pool
)

func hgetRediskeyString(key string, field string) (string, error) {
	redisCon := redisPool.Get()
	defer redisCon.Close()
	result, err := redis.String(redisCon.Do("hget", key, field))
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

// InitAPI func
func InitAPI() {
	martini.Env = martini.Prod
	addr := "redis:" + os.Getenv("NWN_ORDER_REDIS_PORT")
	redisPool = newPool(addr)
	defer redisPool.Close()

	// this is prolly worst practice
	// should disable martini logging
	r := martini.NewRouter()
	mn := martini.New()
	mn.Use(martini.Recovery())
	mn.Use(martini.Static("public"))
	mn.MapTo(r, (*martini.Routes)(nil))
	mn.Action(r.Handle)
	m := &martini.ClassicMartini{mn, r}

	m.Map(redisPool)
	m.Use(render.Renderer())
	m.Get("/", func() string {
		return "1"
	})

	m.Get("/stats", func(r render.Render, pool *redis.Pool, params martini.Params) {
		c := pool.Get()
		defer c.Close()
		k := "order:server"
		v1, err := hgetRediskeyString(k, "BootDate")
		v2, err := hgetRediskeyString(k, "BootTime")
		v3, err := hgetRediskeyString(k, "ModuleName")
		v4, err := hgetRediskeyString(k, "Online")

		if err != nil {
			message := fmt.Sprintf("Could not GET %s", k+":BootTime")
			r.JSON(400, map[string]interface{}{
				"status":  "ERR",
				"message": message})
		} else {
			log.WithFields(log.Fields{"Path": "/stats", "BootDate": v1, "BootTime": v2, "ModuleName": v3, "Online": v4}).Info("Order:API")
			r.JSON(200, map[string]interface{}{
				"status":     "OK",
				"BootDate":   v1,
				"BootTime":   v2,
				"ModuleName": v3,
				"Online":     v4,
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
	runtime.GOMAXPROCS(runtime.NumCPU())

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	go initMain()
	<-done
}
