package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

// config struct
type config struct {
	// core stiff
	RedisPort string `env:"NWN_ORDER_REDIS_PORT" envDefault:"6379"`
	OrderPort string `env:"NWN_ORDER_PORT" envDefault:"5750"`

	// Discord plugin
	PluginDiscord  bool   `env:"NWN_ORDER_PLUGIN_DISCORD_ENABLED" envDefault:"1"`
	DiscordBotKey  string `env:"NWN_ORDER_PLUGIN_DISCORD_BOT_KEY" envDefault:""`
	DiscordBotRoom string `env:"NWN_ORDER_PLUGIN_DISCOD_BOT_ROOM" envDefault:""`

	// Heartbeat plugin
	PluginHearbeat   bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_ENABLED" envDefault:"1"`
	HbVerbose        bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_VERBOSE" envDefault:"false"`
	HbOneMinute      bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_ONE_MINUTE" envDefault:"true"`
	HbFiveMinute     bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_FIVE_MINUTE" envDefault:"true"`
	HbThirtyMinute   bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_THIRTY_MINUTE" envDefault:"true"`
	HbOneHour        bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_ONE_HOUR" envDefault:"true"`
	HbSixHour        bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_SIX_HOUR" envDefault:"true"`
	HbTwelveHour     bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_TWELVE_HOUR" envDefault:"true"`
	HbTwentyfourHour bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_TWENTYFOUR_HOUR" envDefault:"true"`

	// Logs plugin
	PluginLogs bool `env:"NWN_ORDER_PLUGIN_LOG_ENABLED" envDefault:"1"`
}

func initMain() {
	// grab config
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	// start webserver
	go initWebserver()

	// app started
	log.WithFields(log.Fields{"Booted": 1}).Info("Order")

	// connect to redis
	S := 0
	conn, err := net.Dial("udp", "redis:"+cfg.RedisPort)
	for retry := 1; err != nil; retry++ {
		S := strconv.Itoa(retry)
		log.WithFields(log.Fields{"Connected": 0, "Attempt": S}).Warn("Order:Redis")
		if retry > 3 {
			log.WithFields(log.Fields{"Connected": 0, "Attempt": 5}).Fatal("Order:Redis")
		}
		time.Sleep(5 * time.Second)
	}
	log.WithFields(log.Fields{"Connected": 1, "Attempt": S}).Info("Order:Redis")
	conn.Close()

	// start pubsub
	go initPubsub()

	// start plugins
	go initPlugins()

	for {
		time.Sleep(time.Second)
	}
}

func initPlugins() {
	// grab config
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	if cfg.PluginDiscord == true {
		log.WithFields(log.Fields{"Enabled": 1}).Info("Order:Discord")
		go initDiscord()
	} else {
		log.WithFields(log.Fields{"Enabled": 0}).Info("Order:Discord")
	}

	if cfg.PluginHearbeat == true {
		log.WithFields(log.Fields{"Enabled": 1}).Info("Order:Heartbeat")
		go initHeartbeat()
	} else {
		log.WithFields(log.Fields{"Enabled": 0}).Info("Order:Heartbeat")
	}

	if cfg.PluginLogs == true {
		log.WithFields(log.Fields{"Enabled": 1}).Info("Order:Logs")
		go initLog()
	} else {
		log.WithFields(log.Fields{"Enabled": 0}).Info("Order:Logs")
	}
}

func main() {
	done := make(chan bool)
	go initMain()
	<-done // Block forever, not sure if this is best practice.
}
