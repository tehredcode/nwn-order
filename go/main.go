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

// Rds struct
type Rds struct {
	Router *mux.Router
	DB     *redis.Client
}

// Run func
func (a *Rds) Run(addr string) {
	logrus.Fatal(http.ListenAndServe(":8000", a.Router))
}

// InitializeAPI func
func (a *Rds) InitializeAPI() error {
	db := redis.NewClient(&redis.Options{
		Addr: os.Getenv("NWN_ORDER_REDIS_HOST") + ":" + os.Getenv("NWN_ORDER_REDIS_PORT"),
	})
	return nil
}

func (a *Rds) initializeRoutes() {
	a.Router.HandleFunc("/api/status", a.AddStatusHandler).Methods("GET")
}

// AddTodoHandler has access to DB, in your case Redis
// you can replace the steps for Redis.
func (a *Rds) AddTodoHandler() {
	//has access to DB
	a.DB
}








func (a *Rds) initPubsub() {
	p := a.db.PubsubChannel(
		"Discord:Out",
		"Log:Debug",
		"Log:Info",
		"Log:Warning",
		"Log:Fatal",
	)
	for {
		msg, err := p.ReceiveMessage()
		if err != nil {
			logrus.WithFields(logrus.Fields{"Connected": "0"}).Warn("Order:Redis:Pubsub")
		}
		switch msg.Channel {
		case "Discord:Out":
			logrus.WithFields(logrus.Fields{"Pubsub": "1", "Channel": msg.Channel, "Message": msg.Payload}).Info("Order:Pubsub")
		case "Log:Debug":
			logrus.WithFields(logrus.Fields{"Pubsub": "1", "Channel": msg.Channel, "Message": msg.Payload}).Info("Order:Pubsub")
		case "Log:Info":
			logrus.WithFields(logrus.Fields{"Pubsub": "1", "Channel": msg.Channel, "Message": msg.Payload}).Info("Order:Pubsub")
		case "Log:Warning":
			logrus.WithFields(logrus.Fields{"Pubsub": "1", "Channel": msg.Channel, "Message": msg.Payload}).Info("Order:Pubsub")
		case "Log:Fatal":
			logrus.WithFields(logrus.Fields{"Pubsub": "1", "Channel": msg.Channel, "Message": msg.Payload}).Info("Order:Pubsub")
		}
		fmt.Println(msg.Channel, msg.Payload)
	}
}

func initDiscord(a *Rds) {
	logrus.WithFields(logrus.Fields{"BotKey": os.Getenv("NWN_ORDER_PLUGIN_DISCORD_BOT_KEY"), "started": "1"}).Info("Order:Discord")
	discord, err := discordgo.New("Bot " + os.Getenv("NWN_ORDER_PLUGIN_DISCORD_BOT_KEY"))
	errCheck("error creating discord session", err)
	user, err := discord.User("@me")
	errCheck("error retrieving account", err)

	botID = user.ID
	discord.AddHandler(inHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err = discord.UpdateStatus(0, "Order")
		if err != nil {
			logrus.WithFields(logrus.Fields{"Set Status": "0"}).Info("Order:Discord:Error")
		}
		servers := discord.State.Guilds
		logrus.WithFields(logrus.Fields{"Started": 1, "Clients connected": len(servers)}).Info("Order:Discord")
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

	// start pubsub
	go initPubsub(rds)
	logrus.WithFields(logrus.Fields{"Pubsub": 1}).Info("Order")

	// start the web stuff
	//go initHTTP(rds)
	//logrus.WithFields(logrus.Fields{"API": 1}).Info("Order")



	// start plugins
	//go initPlugins(rds)
}

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	go initMain()
	<-done
}
