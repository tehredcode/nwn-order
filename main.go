package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/caarlos0/env"
	"github.com/go-redis/redis"
	"github.com/google/go-github/github"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

var (
	//RedisClient is a var
	RedisClient *redis.Client
)

func startPubsub() {
	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + cfg.RedisPort,
	})
	defer client.Close()

	pubSub := client.Subscribe(
		"heartbeat",
		"input",
		"debug",
		"github",
	)
	for {
		msg, _ := pubSub.ReceiveMessage()
		switch msg.Channel {
		case "heartbeat":

		case "input":

		case "debug":
		}
	}
}

func sendPubsub(pubsubType string, PubsubChannel string, PubsubMessage string) {
	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + cfg.RedisPort,
	})
	defer client.Close()

	if err := client.Publish(PubsubChannel, PubsubMessage).Err(); err != nil {
		panic(err)
	}

	log.WithFields(log.Fields{"Pubsub": pubsubType, "Pubsub Channel": PubsubChannel, "Pubsub Message": PubsubMessage}).Info("Heartbeat")

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

	c := cron.New()
	if cfg.HbOneMinute == true {
		c.AddFunc("@every 1m", func() { sendPubsub("Heartbeat", "heartbeat", "1") })
	}
	log.WithFields(log.Fields{"Heartbeat": "1", "Enabled": cfg.HbOneMinute}).Info("Heartbeat")
	if cfg.HbFiveMinute == true {
		c.AddFunc("@every 5m", func() { sendPubsub("Heartbeat", "heartbeat", "5") })
	}
	log.WithFields(log.Fields{"Heartbeat": "5", "Enabled": cfg.HbFiveMinute}).Info("Heartbeat")
	if cfg.HbThirtyMinute == true {
		c.AddFunc("@every 30m", func() { sendPubsub("Heartbeat", "heartbeat", "30") })
	}
	log.WithFields(log.Fields{"Heartbeat": "30", "Enabled": cfg.HbThirtyMinute}).Info("Heartbeat")
	if cfg.HbOneHour == true {
		c.AddFunc("@every 1h", func() { sendPubsub("Heartbeat", "heartbeat", "60") })
	}
	log.WithFields(log.Fields{"Heartbeat": "60", "Enabled": cfg.HbOneHour}).Info("Heartbeat")
	if cfg.HbSixHour == true {
		c.AddFunc("@every 6h", func() { sendPubsub("Heartbeat", "heartbeat", "360") })
	}
	log.WithFields(log.Fields{"Heartbeat": "360", "Enabled": cfg.HbSixHour}).Info("Heartbeat")
	if cfg.HbTwelveHour == true {
		c.AddFunc("@every 12h", func() { sendPubsub("Heartbeat", "heartbeat", "720") })
	}
	log.WithFields(log.Fields{"Heartbeat": "720", "Enabled": cfg.HbTwelveHour}).Info("Heartbeat")
	if cfg.HbTwentyfourHour == true {
		c.AddFunc("@every 24h", func() { sendPubsub("Heartbeat", "heartbeat", "1440") })
	}
	log.WithFields(log.Fields{"Heartbeat": "1440", "Enabled": cfg.HbTwentyfourHour}).Info("Heartbeat")
	c.Start()

}
