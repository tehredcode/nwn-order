package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/caarlos0/env"
)

func webpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func webserver() {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	http.HandleFunc("/webhook", githubWebhook)
	t := time.Now()
	fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: webserver started :" + cfg.OrderPort + ". webhooks need to be sent to localhost:" + cfg.OrderPort + "/webhook")

	http.HandleFunc("/", webpage)
	log.Fatal(http.ListenAndServe(":"+cfg.OrderPort, nil))
}
