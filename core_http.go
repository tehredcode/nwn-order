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
	data := server{
		Server{ModuleName: Conf.ModuleName},
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
	router := mux.NewRouter()
	router.HandleFunc("/webhook/bitbucket", GithubWebhookHandler)
	router.HandleFunc("/webhook/dockerhub", GithubWebhookHandler)
	router.HandleFunc("/webhook/github", GithubWebhookHandler)
	router.HandleFunc("/webhook/gitlab", GithubWebhookHandler)
	router.HandleFunc("/webhook/gogs", GithubWebhookHandler)
	router.HandleFunc("/api/server", getServerStats).Methods("GET")
	router.HandleFunc("/api/server", setServerStats).Methods("POST")
	http.ListenAndServe(":"+Conf.OrderPort, router)
	log.WithFields(log.Fields{"Port": Conf.OrderPort, "Started": 1}).Fatal("Order:API")
}