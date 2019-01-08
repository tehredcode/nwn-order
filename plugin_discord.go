package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

var (
	commandPrefix string
	botID         string
	botKey        string
)

func initDiscord() {
	// grab config
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	log.WithFields(log.Fields{"BotKey": cfg.DiscordBotKey, "started": "1"}).Info("Order:Discord")
	discord, err := discordgo.New("Bot " + cfg.DiscordBotKey)
	errCheck("error creating discord session", err)
	user, err := discord.User("@me")
	errCheck("error retrieving account", err)

	botID = user.ID
	discord.AddHandler(commandHandler)
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

func errCheck(msg string, err error) {
	if err != nil {
		log.WithFields(log.Fields{"Message": msg, "Error": err}).Fatal("Order:Discord:Error")
	}
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		//Do nothing because the bot is talking
		return
	}

	log.WithFields(log.Fields{"Message Content": message.Content, "Message": message.Message, "Author": message.Author}).Info("Order:Discord:Message")
}
