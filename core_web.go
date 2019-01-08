package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/caarlos0/env"
	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
)

func recieveWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithFields(log.Fields{"Webhook received": "1", "request body": err}).Warn("Github:Commit:Error")
		return
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.WithFields(log.Fields{"Webhook received": err, "Parsed": "0"}).Warn("Github:Commit:Error")
		return
	}

	switch e := event.(type) {

	case *github.PushEvent:
		log.WithFields(log.Fields{"Webhook Received": "Github", "Commit": e.HeadCommit.GetID, "Repo": e.GetRepo(), "Contributor": e.GetPusher()}).Info("Github:Commit")
		go sendPubsub("Github", "Commit", strconv.FormatInt(e.GetPushID(), 10))

	default:
		log.WithFields(log.Fields{"Webhook Received": "Github", "Event": github.WebHookType(r), "Supported": "0"}).Warn("Github:Commit")
		return
	}
}

func initWebserver() {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	http.HandleFunc("/webhook", recieveWebhook)
	log.WithFields(log.Fields{"Started": 1, "Port": cfg.OrderPort, "path": "/webhook"}).Info("Order:Webserver")

	http.ListenAndServe(":"+cfg.OrderPort, nil)
}
