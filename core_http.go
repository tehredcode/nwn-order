package main

import (
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
	c := Config{}
	data := server{
		Server{ModuleName: c.ModuleName},
		Server{BootTime: ModuleBootTime()},
		Server{BootDate: ModuleBootDate()},
		Server{Players: ModulePlayers()},
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
	r := mux.NewRouter()
	r.HandleFunc("/webhook/dockerhub", DockerhubWebhookHandler)
	r.HandleFunc("/webhook/github", GithubWebhookHandler)
	r.HandleFunc("/webhook/gitlab", GitlabWebhookHandler)
	r.HandleFunc("/api/server", getServerStats)
	http.ListenAndServe(":"+c.OrderPort, r)
	log.WithFields(log.Fields{"Port": c.OrderPort, "Started": 1}).Info("Order:API")
}
