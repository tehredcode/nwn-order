package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var (
	commandPrefix string
	botID         string
	botKey        string
)

func initDiscord() {
	botkey := os.Getenv("DiscordBotKey")
	log.WithFields(log.Fields{"BotKey": botkey}).Info("Order:Discord:Status")
	discord, err := discordgo.New("Bot " + botkey)
	errCheck("error creating discord session", err)
	user, err := discord.User("@me")
	errCheck("error retrieving account", err)

	botID = user.ID
	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err = discord.UpdateStatus(0, "Order is coming")
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

	content := message.Content
	fmt.Printf(content)
	fmt.Printf("Message: %+v || From: %s\n", message.Message, message.Author)
}
