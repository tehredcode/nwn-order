package heartbeat

import (
	"os"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

func initHeartbeat() {
	// app started
	log.WithFields(log.Fields{"Started": 1}).Info("Order:Heartbeat")

	c := cron.New()
	if os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_ONE_MINUTE") == "1" {
		c.AddFunc("@every 1m", func() { sendPubsub("Heartbeat", "Heartbeat", "1") })
	}
	log.WithFields(log.Fields{"Heartbeat": "1", "Enabled": os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_ONE_MINUTE")}).Info("Heartbeat")

	if os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_FIVE_MINUTE") == "1" {
		c.AddFunc("@every 5m", func() { sendPubsub("Heartbeat", "Heartbeat", "5") })
	}
	log.WithFields(log.Fields{"Heartbeat": "5", "Enabled": os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_FIVE_MINUTE")}).Info("Heartbeat")

	if os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_THIRTY_MINUTE") == "1" {
		c.AddFunc("@every 30m", func() { sendPubsub("Heartbeat", "Heartbeat", "30") })
	}
	log.WithFields(log.Fields{"Heartbeat": "30", "Enabled": os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_THIRTY_MINUTE")}).Info("Heartbeat")

	if os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_ONE_HOUR") == "1" {
		c.AddFunc("@every 1h", func() { sendPubsub("Heartbeat", "Heartbeat", "60") })
	}
	log.WithFields(log.Fields{"Heartbeat": "60", "Enabled": os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_ONE_HOUR")}).Info("Heartbeat")

	if os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_SIX_HOUR") == "1" {
		c.AddFunc("@every 6h", func() { sendPubsub("Heartbeat", "Heartbeat", "360") })
	}
	log.WithFields(log.Fields{"Heartbeat": "360", "Enabled": os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_SIX_HOUR")}).Info("Heartbeat")

	if os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_TWELVE_HOUR") == "1" {
		c.AddFunc("@every 12h", func() { sendPubsub("Heartbeat", "Heartbeat", "720") })
	}
	log.WithFields(log.Fields{"Heartbeat": "720", "Enabled": os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_TWELVE_HOUR")}).Info("Heartbeat")

	if os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_TWENTYFOUR_HOUR") == "1" {
		c.AddFunc("@every 24h", func() { sendPubsub("Heartbeat", "Heartbeat", "1440") })
	}
	log.WithFields(log.Fields{"Heartbeat": "1440", "Enabled": os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_TWENTYFOUR_HOUR")}).Info("Heartbeat")
	c.Start()
}
