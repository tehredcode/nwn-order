package main

import (
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

func initHeartbeat() {
	// app started
	conf := Config{}
	log.WithFields(log.Fields{"Started": 1}).Info("Order:Heartbeat")

	c := cron.New()
	if conf.HbOneMinute == true {
		c.AddFunc("@every 1m", func() { sendPubsub("Heartbeat", "Heartbeat", "1") })
	}
	log.WithFields(log.Fields{"Heartbeat": "1", "Enabled": conf.HbOneMinute}).Info("Heartbeat")
	if conf.HbFiveMinute == true {
		c.AddFunc("@every 5m", func() { sendPubsub("Heartbeat", "Heartbeat", "5") })
	}
	log.WithFields(log.Fields{"Heartbeat": "5", "Enabled": conf.HbFiveMinute}).Info("Heartbeat")
	if conf.HbThirtyMinute == true {
		c.AddFunc("@every 30m", func() { sendPubsub("Heartbeat", "Heartbeat", "30") })
	}
	log.WithFields(log.Fields{"Heartbeat": "30", "Enabled": conf.HbThirtyMinute}).Info("Heartbeat")
	if conf.HbOneHour == true {
		c.AddFunc("@every 1h", func() { sendPubsub("Heartbeat", "Heartbeat", "60") })
	}
	log.WithFields(log.Fields{"Heartbeat": "60", "Enabled": conf.HbOneHour}).Info("Heartbeat")
	if conf.HbSixHour == true {
		c.AddFunc("@every 6h", func() { sendPubsub("Heartbeat", "Heartbeat", "360") })
	}
	log.WithFields(log.Fields{"Heartbeat": "360", "Enabled": conf.HbSixHour}).Info("Heartbeat")
	if conf.HbTwelveHour == true {
		c.AddFunc("@every 12h", func() { sendPubsub("Heartbeat", "Heartbeat", "720") })
	}
	log.WithFields(log.Fields{"Heartbeat": "720", "Enabled": conf.HbTwelveHour}).Info("Heartbeat")
	if conf.HbTwentyfourHour == true {
		c.AddFunc("@every 24h", func() { sendPubsub("Heartbeat", "Heartbeat", "1440") })
	}
	log.WithFields(log.Fields{"Heartbeat": "1440", "Enabled": conf.HbTwentyfourHour}).Info("Heartbeat")
	c.Start()
}
