package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

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

func sendPubsub(LogMessage string, PubsubChannel string, PubsubMessage string) {
	r := redis.NewClient(&redis.Options{Addr: os.Getenv("NWN_ORDER_REDIS_HOST") + ":" + os.Getenv("NWN_ORDER_REDIS_PORT")})
	err := r.Publish(PubsubChannel, PubsubMessage).Err()
	if err != nil {
		log.WithFields(log.Fields{"Channel": PubsubChannel, "Message": PubsubMessage}).Error(LogMessage)
		panic(err)
	}
	log.WithFields(log.Fields{"Channel": PubsubChannel, "Message": PubsubMessage}).Info(LogMessage)
}
