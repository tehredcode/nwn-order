package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var (
	commandPrefix string
	botID         string
	botKey        string
)

func mapSubexpNames(m, n []string) map[string]string {
	m, n = m[1:], n[1:]
	r := make(map[string]string, len(m))
	for i := range n {
		r[n[i]] = m[i]
	}
	return r
}

func initDiscord() {
	log.WithFields(log.Fields{"BotKey": os.Getenv("NWN_ORDER_DISCORD_BOT_KEY"), "started": "1"}).Info("Order:Discord")
	discord, err := discordgo.New("Bot " + os.Getenv("NWN_ORDER_DISCORD_BOT_KEY"))
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

func errCheck(msg string, err error) {
	if err != nil {
		log.WithFields(log.Fields{"Message": msg, "Error": err}).Fatal("Order:Discord:Error")
	}
}

func replyHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		//Do nothing because the bot is talking
		return
	}

	if message.ChannelID == os.Getenv("NWN_ORDER_DISCORD_BOT_ROOM") {
		sendPubsub(message.ChannelID, "Discord:Out", "["+message.Author.Username+"] "+message.Content)
		log.WithFields(log.Fields{"Message Content": message.Content, "Message": message.Message, "Author": message.Author}).Info("Order:Discord:Message")
		return
	}
}

func inHandler(discord *discordgo.Session, m string) {
	discord.ChannelMessageSend(os.Getenv("NWN_ORDER_DISCORD_BOT_ROOM"), m)
}
