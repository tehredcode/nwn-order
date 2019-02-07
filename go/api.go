package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// apiRoot
type apiRootStruct struct {
	Online string
}

// api root
func apiRoot(w http.ResponseWriter, r *http.Request) {
	stats := apiRootStruct{"1"}
	js, err := json.Marshal(stats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// ServerStats struct
type ServerStats struct {
	BootDate   string
	BootTime   string
	ModuleName string
	Online     string
}

// api stats
func apiStats(w http.ResponseWriter, r *http.Request) {
	k := "order:server"
	v1, err := hgetRediskeyString(k, "BootDate")
	v2, err := hgetRediskeyString(k, "BootTime")
	v3, err := hgetRediskeyString(k, "ModuleName")
	v4, err := hgetRediskeyString(k, "Online")

	stats := ServerStats{v1, v2, v3, v4}

	js, err := json.Marshal(stats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// InitAPI func
func InitAPI() {
	log.Println("server started")

	m := mux.NewRouter()

	m.HandleFunc("/", apiRoot)
	m.HandleFunc("/api/stats", apiStats)

	if os.Getenv("NWN_ORDER_WEBHOOKS") == "1" {
		log.Println("Webhooks started: " + os.Getenv("NWN_ORDER_ORDER_PORT"))
		m.HandleFunc("/webhook/Dockerhub", DockerhubWebhookHandler)
		m.HandleFunc("/webhook/Github", GithubWebhookHandler)
		m.HandleFunc("/webhook/Gitlab", GitlabWebhookHandler)
	}

	addr := "redis:" + os.Getenv("NWN_ORDER_REDIS_PORT")
	RedisPool = newPool(addr)
	defer RedisPool.Close()

	log.Println("API started: " + os.Getenv("NWN_ORDER_PORT"))
	http.ListenAndServe(":"+os.Getenv("NWN_ORDER_PORT"), m)
}
