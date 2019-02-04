package discord

import (
	"os"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	rds "github.com/urothis/nwn-order/go/redis"
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

	if message.ChannelID == os.Getenv("NWN_ORDER_PLUGIN_DISCOD_BOT_ROOM") {
		rds.SendPubsub(message.ChannelID, "Discord:Out", "["+message.Author.Username+"] "+message.Content)
		log.WithFields(log.Fields{"Message Content": message.Content, "Message": message.Message, "Author": message.Author}).Info("Order:Discord:Message")
		return
	}
}

func inHandler(discord *discordgo.Session, m string) {
	discord.ChannelMessageSend(os.Getenv("NWN_ORDER_PLUGIN_DISCOD_BOT_ROOM"), m)
}
