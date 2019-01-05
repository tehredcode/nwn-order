package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/caarlos0/env"
	"github.com/google/go-github/github"
	"github.com/jasonlvhit/gocron"
)

func heartbeatWebhook(ticker string, verbose bool) {
	t := time.Now()
	msg := ("O [" + t.Format("15:04:05") + "] [NWN_Order] Pubsub Event: channel=heartbeat message=" + ticker)
	sendPubsub(msg, "heartbeat", ticker)
}

func githubWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error validating request body: err=%s\n", err)
		return
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		fmt.Printf("could not parse webhook: err=%s\n", err)
		return
	}

	switch e := event.(type) {

	case *github.PushEvent:
		msg := ("O [" + time.Now().Format("15:04:05") + "] [NWN_Order] Webhook Event: channel=innwserver message=repoupdate | " + *e.Sender.Login + " made a commit to module repo")
		go sendPubsub(msg, "github", "commit")

	default:
		fmt.Printf("Only push events supported, unknown webhook event type %s\n", github.WebHookType(r))
		return
	}
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.WithFields(log.Fields{
		"boot": "started",
	}).Info("Order has Started")

	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	conn, err := net.Dial("udp", "redis:"+cfg.RedisPort)
	for retry := 1; err != nil; retry++ {
		s := strconv.Itoa(retry)
		log.WithFields(log.Fields{
			"Connected": "0",
			"Attempt":   s,
		}).Warn("Redis not connected")

		if retry > 4 {
			log.WithFields(log.Fields{
				"Redis": "0",
			}).Fatal("Redis not connected")
			os.Exit(1)
		}
		time.Sleep(5 * time.Second)
	}
	conn.Close()

	log.WithFields(log.Fields{
		"Redis": "1",
	}).Info("Redis connected")

	// start pubsub
	go startPubsub()
	log.WithFields(log.Fields{
		"Pubsub": "1",
	}).Info("Pubsub started")

	// start webhook reciever
	go webserver()
	log.WithFields(log.Fields{
		"Webserver": "1",
	}).Info("Webserver started")

	// start the heartbeat timers
	if cfg.HbOneMinute == true {
		gocron.Every(1).Minute().Do(heartbeatWebhook, "1")
	}
	log.WithFields(log.Fields{"Heartbeat": "1", "Enabled": cfg.HbOneMinute}).Info("Heartbeat")

	if cfg.HbFiveMinute == true {
		gocron.Every(5).Minutes().Do(heartbeatWebhook, "5")
	}
	log.WithFields(log.Fields{"Heartbeat": "5", "Enabled": cfg.HbFiveMinute}).Info("Heartbeat")

	if cfg.HbThirtyMinute == true {
		gocron.Every(30).Minutes().Do(heartbeatWebhook, "30")
	}
	log.WithFields(log.Fields{"Heartbeat": "30", "Enabled": cfg.HbThirtyMinute}).Info("Heartbeat")

	if cfg.HbOneHour == true {
		gocron.Every(1).Hour().Do(heartbeatWebhook, "60")
	}
	log.WithFields(log.Fields{"Heartbeat": "360", "Enabled": cfg.HbOneHour}).Info("Heartbeat")

	if cfg.HbSixHour == true {
		gocron.Every(6).Hours().Do(heartbeatWebhook, "360")
	}
	log.WithFields(log.Fields{"Heartbeat": "360", "Enabled": cfg.HbSixHour}).Info("Heartbeat")

	if cfg.HbTwelveHour == true {
		gocron.Every(12).Hours().Do(heartbeatWebhook, "720")
	}
	log.WithFields(log.Fields{"Heartbeat": "720", "Enabled": cfg.HbTwelveHour}).Info("Heartbeat")

	if cfg.HbTwentyfourHour == true {
		gocron.Every(24).Hours().Do(heartbeatWebhook, "1440")
	}
	log.WithFields(log.Fields{"Heartbeat": "1440", "Enabled": cfg.HbTwentyfourHour}).Info("Heartbeat")

	<-gocron.Start()
}
