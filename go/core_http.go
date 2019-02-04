package main

import (
	"fmt"
	"os"

	"encoding/json"
	"net/http"

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
	rds2 := NewRedisClient(os.Getenv("NWN_ORDER_REDIS_HOST")+":"+os.Getenv("NWN_ORDER_REDIS_PORT"), "")
	rkey := os.Getenv("NWN_ORDER_PORT") + ":server"
	value, _ := rds2.HMGet(rkey,
		"BootTime",
		"BootDate",
		"Online",
	)
	fmt.Printf("%#v", value)

	//rds := redis.NewClient(&redis.Options{Addr: os.Getenv("NWN_ORDER_REDIS_HOST") + ":" + os.Getenv("NWN_ORDER_REDIS_PORT")})
	//s, _ := rds.HMGet(rkey, "BootTime", "BootDate", "Online").Result()
	//fmt.Printf("%#v", s)

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

func setServerStats(w http.ResponseWriter, r *http.Request) {
}

func initHTTP() {
	r := mux.NewRouter()
	r.HandleFunc("/webhook/dockerhub", DockerhubWebhookHandler)
	r.HandleFunc("/webhook/github", GithubWebhookHandler)
	r.HandleFunc("/webhook/gitlab", GitlabWebhookHandler)
	r.HandleFunc("/api/server", getServerStats)
	http.ListenAndServe(":"+os.Getenv("NWN_ORDER_PORT"), r)
	log.WithFields(log.Fields{"Port": os.Getenv("NWN_ORDER_PORT"), "Started": 1}).Info("Order:API")
}
