package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	rds "github.com/urothis/nwn-order/go/redis"
)

// Server struct
type Server struct {
	ModuleName string `json:"modulename,omitempty"`
	BootTime   string `json:"boottime,omitempty"`
	BootDate   string `json:"bootdate,omitempty"`
	Players    string `json:"players,omitempty"`
}

type server []Server

func getServerStats(c *rds.Client, w http.ResponseWriter, r *http.Request) {
	rkey := os.Getenv("NWN_ORDER_PORT") + ":server"
	value, _ := c.HMGet(rkey,
		"BootTime",
		"BootDate",
		"Online",
	)
	fmt.Printf("%#v", value)

	data := server{
		Server{ModuleName: os.Getenv("NWN_ORDER_MODULE_NAME")},
		//Server{BootTime: s[1]},
		//Server{BootDate: s[2]},
		//Server{Players: s[3]},
	}
	json.NewEncoder(w).Encode(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)

}

func redisHandler(c *redis.Client,
	f func(c *redis.Client, w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { f(c, w, r) })
}
