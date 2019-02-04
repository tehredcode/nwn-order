package main

import (
	config "github.com/Urothis-nwn-Order/nwn-order/blob/dev/go/config"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Server struct
type Server struct {
	ModuleName string `json:"modulename,omitempty"`
	BootTime   string `json:"boottime,omitempty"`
	BootDate   string `json:"bootdate,omitempty"`
	Players    string `json:"players,omitempty"`
}

type server []Server

func getServerStats(w http.ResponseWriter, r *http.Request) {
	c := Config{}
	rds := redis.NewClient(&redis.Options{Addr: "redis:" + config.RedisPort})
	rkey := c.ModuleName + ":server"
	ModuleBootTime, _ := rds.HGet(rkey, "BootTime").Result()
	rkey = c.ModuleName + ":server"
	ModuleBootDate, _ := rds.HGet(rkey, "BootDate").Result()
	rkey = c.ModuleName + ":server"
	ModulePlayers, _ := rds.HGet(rkey, "Online").Result()

	data := server{
		Server{ModuleName: c.ModuleName},
		Server{BootTime: ModuleBootTime},
		Server{BootDate: ModuleBootDate},
		Server{Players: ModulePlayers},
	}
	json.NewEncoder(w).Encode(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func setServerStats(w http.ResponseWriter, r *http.Request) {
}

func initHTTP() {
	c := Config{}
	err := env.Parse(&c)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/webhook/dockerhub", DockerhubWebhookHandler)
	r.HandleFunc("/webhook/github", GithubWebhookHandler)
	r.HandleFunc("/webhook/gitlab", GitlabWebhookHandler)
	r.HandleFunc("/api/server", getServerStats)
	http.ListenAndServe(":"+c.OrderPort, r)
	log.WithFields(log.Fields{"Port": c.OrderPort, "Started": 1}).Info("Order:API")
}
