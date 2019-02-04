package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/urothis/nwn-order/go/redis"
)

// Server struct
type Server struct {
	ModuleName string `json:"modulename,omitempty"`
	BootTime   string `json:"boottime,omitempty"`
	BootDate   string `json:"bootdate,omitempty"`
	Players    string `json:"players,omitempty"`
}

type server []Server

func getServerStats(c *RedisClient, w http.ResponseWriter, r *http.Request) {
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

func initHTTP() {
	//Initialize Redis Client
	client := redis.InitClient()
	//Get current redis instance to get passed to different Gorilla-Mux Handlers
	redisHandler := &RedisInstance{RInstance: &client}

	r := mux.NewRouter()
	r.HandleFunc("/webhook/dockerhub", DockerhubWebhookHandler)
	r.HandleFunc("/webhook/github", GithubWebhookHandler)
	r.HandleFunc("/webhook/gitlab", GitlabWebhookHandler)

	r.HandleFunc("/api/server", redisHandler.AddTodoHandler).Methods("POST")

	http.ListenAndServe(":"+os.Getenv("NWN_ORDER_PORT"), r)
	log.WithFields(log.Fields{"Port": os.Getenv("NWN_ORDER_PORT"), "Started": 1}).Info("Order:API")
}
