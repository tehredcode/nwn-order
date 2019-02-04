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
	"github.com/urothis/nwn-order/go/log"
)

func initHTTP(c *RedisInstance) {
	r := mux.NewRouter()
	r.HandleFunc("/webhook/dockerhub", DockerhubWebhookHandler)
	r.HandleFunc("/webhook/github", GithubWebhookHandler)
	r.HandleFunc("/webhook/gitlab", GitlabWebhookHandler)
	r.HandleFunc("/api/server", redisHandler(c, getServerStats(c))).Methods("POST")

	http.ListenAndServe(":"+os.Getenv("NWN_ORDER_PORT"), r)
	log.WithFields(log.Fields{"Port": os.Getenv("NWN_ORDER_PORT"), "Started": 1}).Info("Order:API")
}

func initPubsub() {
	addr := os.Getenv("NWN_ORDER_REDIS_HOST") + ":" + os.Getenv("NWN_ORDER_REDIS_PORT")
	client := InitClient()
	r := &Client{RInstance: &client}

	defer r.Close()

	pubSub, err := r.Subscribe()
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

func initDiscord() {
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

func initMain() {
	// app started
	logrus.WithFields(logrus.Fields{"Booted": 1}).Info("Order")

	// grab redis client
	client := InitRedisClient()
	redisHandler := &RedisInstance{RInstance: &client}

	// start the web stuff
	go initHTTP()
	logrus.WithFields(logrus.Fields{"API": 1}).Info("Order")

	// start pubsub
	go initPubsub()
	logrus.WithFields(logrus.Fields{"Pubsub": 1}).Info("Order")

	// start plugins
	go initPlugins()
}

func initPlugins() {
	if os.Getenv("NWN_ORDER_PLUGIN_DISCORD_ENABLED") == "1" {
		logrus.WithFields(logrus.Fields{"Enabled": 1}).Info("Order:Discord")
		go initDiscord()
	} else {
		logrus.WithFields(logrus.Fields{"Enabled": 0}).Info("Order:Discord")
	}

	if os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_ENABLED") == "1" {
		logrus.WithFields(logrus.Fields{"Enabled": 1}).Info("Order:Heartbeat")
		go initHeartbeat()
	} else {
		logrus.WithFields(logrus.Fields{"Enabled": 0}).Info("Order:Heartbeat")
	}

	if os.Getenv("NWN_ORDER_PLUGIN_LOG_ENABLED") == "1" {
		logrus.WithFields(logrus.Fields{"Enabled": 1}).Info("Order:log")
		go initLog()
	} else {
		logrus.WithFields(logrus.Fields{"Enabled": 0}).Info("Order:log")
	}
}

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	go initMain()
	<-done
}
