package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
	logrus "github.com/sirupsen/logrus"
	api "github.com/urothis/nwn-order/go/api"
	log "github.com/urothis/nwn-order/go/log"
	rds "github.com/urothis/nwn-order/go/redis"
)

func initHTTP(c *rds.Client) {
	r := mux.NewRouter()

	r.HandleFunc("/webhook/dockerhub", api.DockerhubWebhookHandler)
	r.HandleFunc("/webhook/github", api.GithubWebhookHandler)
	r.HandleFunc("/webhook/gitlab", api.GitlabWebhookHandler)
	r.HandleFunc("/api/server", rds.redisHandler(client, api.GetServerStats)).Methods("POST")

	http.ListenAndServe(":"+os.Getenv("NWN_ORDER_PORT"), r)
	logrus.WithFields(logrus.Fields{"Port": os.Getenv("NWN_ORDER_PORT"), "Started": 1}).Info("Order:API")
}

func initPubsub(c *rds.Client) {
	pubSub, err := c.Subscribe()
	err = pubSub.Subscribe(
		"Discord:Out",
		"Log:Debug",
		"Log:Info",
		"Log:Warning",
		"Log:Fatal",
	)
	if err != nil {
		log.WithFields(log.Fields{"Connected": "0", "Please confirm redis is connected": "1"}).Fatal("Order:Redis")
	}
	defer pubSub.Close()
	log.WithFields(log.Fields{"Discord:Out": "1", "Log:Debug": "1", "Log:Info": "1", "Log:Warning": "1", "Log:Fatal": "1"}).Info("Order:Redis:Pubsub:Subscribe")

	for {
		msg, err := pubSub.ReceiveMessage()
		if err != nil {
			log.WithFields(log.Fields{"Connected": "0"}).Warn("Order:Redis:Pubsub")
		}
		switch msg.Channel {
		case "Discord:Out":
			log.WithFields(log.Fields{"Pubsub": "1", "Channel": msg.Channel, "Message": msg.Payload}).Info("Order:Pubsub")
		case "Log:Debug":
			log.WithFields(log.Fields{"Pubsub": "1", "Channel": msg.Channel, "Message": msg.Payload}).Info("Order:Pubsub")
		case "Log:Info":
			log.WithFields(log.Fields{"Pubsub": "1", "Channel": msg.Channel, "Message": msg.Payload}).Info("Order:Pubsub")
		case "Log:Warning":
			log.WithFields(log.Fields{"Pubsub": "1", "Channel": msg.Channel, "Message": msg.Payload}).Info("Order:Pubsub")
		case "Log:Fatal":
			log.WithFields(log.Fields{"Pubsub": "1", "Channel": msg.Channel, "Message": msg.Payload}).Info("Order:Pubsub")
		}
		fmt.Println(msg.Channel, msg.Payload)
	}
}

func initDiscord(c *rds.Client) {
	log.WithFields(log.Fields{"BotKey": os.Getenv("NWN_ORDER_PLUGIN_DISCORD_BOT_KEY"), "started": "1"}).Info("Order:Discord")
	discord, err := discordgo.New("Bot " + os.Getenv("NWN_ORDER_PLUGIN_DISCORD_BOT_KEY"))
	errCheck("error creating discord session", err)
	user, err := discord.User("@me")
	errCheck("error retrieving account", err)

	botID = user.ID
	discord.AddHandler(inHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err = discord.UpdateStatus(0, "Order")
		if err != nil {
			log.WithFields(log.Fields{"Set Status": "0"}).Info("Order:Discord:Error")
		}
		servers := discord.State.Guilds
		log.WithFields(log.Fields{"Started": 1, "Clients connected": len(servers)}).Info("Order:Discord")
	})

	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	commandPrefix = "!"

	<-make(chan struct{})
}

func initPlugins(rds *RedisInstance) {
	if os.Getenv("NWN_ORDER_PLUGIN_DISCORD_ENABLED") == "1" {
		logrus.WithFields(logrus.Fields{"Enabled": 1}).Info("Order:Discord")
		go initDiscord(rds)
	} else {
		logrus.WithFields(logrus.Fields{"Enabled": 0}).Info("Order:Discord")
	}

	if os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_ENABLED") == "1" {
		logrus.WithFields(logrus.Fields{"Enabled": 1}).Info("Order:Heartbeat")
		go initHeartbeat(rds)
	} else {
		logrus.WithFields(logrus.Fields{"Enabled": 0}).Info("Order:Heartbeat")
	}

	if os.Getenv("NWN_ORDER_PLUGIN_LOG_ENABLED") == "1" {
		logrus.WithFields(logrus.Fields{"Enabled": 1}).Info("Order:log")
		go initLog(rds)
	} else {
		logrus.WithFields(logrus.Fields{"Enabled": 0}).Info("Order:log")
	}
}

func initMain() {
	// app started
	logrus.WithFields(logrus.Fields{"Booted": 1}).Info("Order")

	// grab redis client
	client := InitRedisClient()
	rds := &RedisInstance{rds.Client: &client}

	// start the web stuff
	go initHTTP(rds)
	logrus.WithFields(logrus.Fields{"API": 1}).Info("Order")

	// start pubsub
	go initPubsub(rds)
	logrus.WithFields(logrus.Fields{"Pubsub": 1}).Info("Order")

	// start plugins
	go initPlugins(rds)
}

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	go initMain()
	<-done
}
