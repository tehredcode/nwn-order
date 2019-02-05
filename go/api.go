package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// ServerStats struct
type ServerStats struct {
	ModuleName string `json:"modulename,omitempty"`
	BootTime   string `json:"boottime,omitempty"`
	BootDate   string `json:"bootdate,omitempty"`
	Players    string `json:"players,omitempty"`
}

type server []ServerStats

// GetServerStats func
func GetServerStats(a *order.Rds, w http.ResponseWriter, r *http.Request) {
	rkey := os.Getenv("NWN_ORDER_PORT") + ":server"
	value, _ := a.HMGet(rkey,
		"BootTime",
		"BootDate",
		"Online",
	)
	fmt.Printf("%#v", value)

	data := server{
		ServerStats{ModuleName: os.Getenv("NWN_ORDER_MODULE_NAME")},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
