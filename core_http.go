package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func initWebhook() {
	http.HandleFunc("/webhook", webhookHandler)
	log.WithFields(log.Fields{"Started": 1, "Port": Conf.OrderPort, "path": "/webhook"}).Info("Order:Webserver")

	http.ListenAndServe(":"+Conf.OrderPort, nil)
}

func initHTTP() {
	initWebhook()
}
