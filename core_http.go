package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Server struct
type Server struct {
	ModuleName string
	BootTime   string
	BootDate   string
	Players    string
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
	router.HandleFunc("/webhook", webhookHandler)
	router.HandleFunc("/api/server", getServerStats).Methods("GET")
	router.HandleFunc("/api/server", setServerStats).Methods("POST")
	http.ListenAndServe(":"+Conf.OrderPort, router)
}
