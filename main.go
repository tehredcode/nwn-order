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

type config struct {
	RedisPort        string `env:"NWN_ORDER_REDIS_PORT" envDefault:"6379"`
	OrderPort        string `env:"NWN_ORDER_PORT" envDefault:"5750"`
	DiscordBotKey    string `env:"NWN_ORDER_DISCORD_BOT_KEY" envDefault:"tiddies"`
	HbVerbose        bool   `env:"NWN_ORDER_HB_VERBOSE" envDefault:"false"`
	HbOneMinute      bool   `env:"NWN_ORDER_HB_ONE_MINUTE" envDefault:"true"`
	HbFiveMinute     bool   `env:"NWN_ORDER_HB_FIVE_MINUTE" envDefault:"true"`
	HbThirtyMinute   bool   `env:"NWN_ORDER_HB_THIRTY_MINUTE" envDefault:"true"`
	HbOneHour        bool   `env:"NWN_ORDER_HB_ONE_HOUR" envDefault:"true"`
	HbSixHour        bool   `env:"NWN_ORDER_HB_SIX_HOUR" envDefault:"true"`
	HbTwelveHour     bool   `env:"NWN_ORDER_HB_TWELVE_HOUR" envDefault:"true"`
	HbTwentyfourHour bool   `env:"NWN_ORDER_HB_TWENTYFOUR_HOUR" envDefault:"true"`
}

var (
	//RedisClient is a var
	RedisClient *redis.Client
)

func initMain() {
	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	log.WithFields(log.Fields{"Booted": 1}).Info("Order")

	go initDiscord()

	S := 0
	conn, err := net.Dial("udp", "redis:6379")
	for retry := 1; err != nil; retry++ {
		S := strconv.Itoa(retry)
		log.WithFields(log.Fields{"Connected": 0, "Attempt": S}).Warn("Order:Redis")
		if retry > 4 {
			log.WithFields(log.Fields{"Connected": 0, "Attempt": S}).Fatal("Order:Redis")
		}
		time.Sleep(5 * time.Second)
	}
	log.WithFields(log.Fields{"Connected": 1, "Attempt": S}).Info("Order:Redis")
	conn.Close()

	// start pubsub
	go initPubsub()

	go webserver()

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

	for {
		time.Sleep(time.Second)
	}
}

func initPubsub() {
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
	log.WithFields(log.Fields{"heartbeat": "1", "input": "1", "debug": "1", "github": "1"}).Info("Order:Redis:Pubsub:Subscribe")
	for {
		msg, _ := pubSub.ReceiveMessage()
		switch msg.Channel {
		case "heartbeat":

		case "input":

		case "debug":
		}
	}
}

func sendPubsub(LogMessage string, PubsubChannel string, PubsubMessage string) {
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

	log.WithFields(log.Fields{"Channel": PubsubChannel, "Message": PubsubMessage}).Info(LogMessage)
}

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

func webserver() {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	http.HandleFunc("/webhook", recieveWebhook)
	log.WithFields(log.Fields{"Started": 1, "Port": cfg.OrderPort, "path": "/webhook"}).Info("Order:Webserver")

	http.ListenAndServe(":"+cfg.OrderPort, nil)
}

func main() {
	done := make(chan bool)
	go initMain()
	<-done // Block forever, not sure if this is best practice.
}
