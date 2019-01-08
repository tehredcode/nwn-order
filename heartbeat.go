package main

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

func initHeartbeat() {
	// grab config
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	// app started
	log.WithFields(log.Fields{"Started": 1}).Info("Order:Heartbeat")

	c := cron.New()
	if cfg.HbOneMinute == true {
		c.AddFunc("@every 1m", func() { sendPubsub("Heartbeat", "heartbeat", "1") })
	}
	log.WithFields(log.Fields{"Heartbeat": "1", "Enabled": cfg.HbOneMinute}).Info("Heartbeat")
	if cfg.HbFiveMinute == true {
		c.AddFunc("@every 5m", func() { sendPubsub("Heartbeat", "heartbeat", "5") })
	}
	log.WithFields(log.Fields{"Heartbeat": "5", "Enabled": cfg.HbFiveMinute}).Info("Heartbeat")
	if cfg.HbThirtyMinute == true {
		c.AddFunc("@every 30m", func() { sendPubsub("Heartbeat", "heartbeat", "30") })
	}
	log.WithFields(log.Fields{"Heartbeat": "30", "Enabled": cfg.HbThirtyMinute}).Info("Heartbeat")
	if cfg.HbOneHour == true {
		c.AddFunc("@every 1h", func() { sendPubsub("Heartbeat", "heartbeat", "60") })
	}
	log.WithFields(log.Fields{"Heartbeat": "60", "Enabled": cfg.HbOneHour}).Info("Heartbeat")
	if cfg.HbSixHour == true {
		c.AddFunc("@every 6h", func() { sendPubsub("Heartbeat", "heartbeat", "360") })
	}
	log.WithFields(log.Fields{"Heartbeat": "360", "Enabled": cfg.HbSixHour}).Info("Heartbeat")
	if cfg.HbTwelveHour == true {
		c.AddFunc("@every 12h", func() { sendPubsub("Heartbeat", "heartbeat", "720") })
	}
	log.WithFields(log.Fields{"Heartbeat": "720", "Enabled": cfg.HbTwelveHour}).Info("Heartbeat")
	if cfg.HbTwentyfourHour == true {
		c.AddFunc("@every 24h", func() { sendPubsub("Heartbeat", "heartbeat", "1440") })
	}
	log.WithFields(log.Fields{"Heartbeat": "1440", "Enabled": cfg.HbTwentyfourHour}).Info("Heartbeat")
	c.Start()
}
