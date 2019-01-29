package main

import (
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

func initHeartbeat() {
	// app started
	log.WithFields(log.Fields{"Started": 1}).Info("Order:Heartbeat")

	c := cron.New()
	if Conf.HbOneMinute == true {
		c.AddFunc("@every 1m", func() { sendPubsub("Heartbeat", "heartbeat", "1") })
	}
	log.WithFields(log.Fields{"Heartbeat": "1", "Enabled": Conf.HbOneMinute}).Info("Heartbeat")
	if Conf.HbFiveMinute == true {
		c.AddFunc("@every 5m", func() { sendPubsub("Heartbeat", "heartbeat", "5") })
	}
	log.WithFields(log.Fields{"Heartbeat": "5", "Enabled": Conf.HbFiveMinute}).Info("Heartbeat")
	if Conf.HbThirtyMinute == true {
		c.AddFunc("@every 30m", func() { sendPubsub("Heartbeat", "heartbeat", "30") })
	}
	log.WithFields(log.Fields{"Heartbeat": "30", "Enabled": Conf.HbThirtyMinute}).Info("Heartbeat")
	if Conf.HbOneHour == true {
		c.AddFunc("@every 1h", func() { sendPubsub("Heartbeat", "heartbeat", "60") })
	}
	log.WithFields(log.Fields{"Heartbeat": "60", "Enabled": Conf.HbOneHour}).Info("Heartbeat")
	if Conf.HbSixHour == true {
		c.AddFunc("@every 6h", func() { sendPubsub("Heartbeat", "heartbeat", "360") })
	}
	log.WithFields(log.Fields{"Heartbeat": "360", "Enabled": Conf.HbSixHour}).Info("Heartbeat")
	if Conf.HbTwelveHour == true {
		c.AddFunc("@every 12h", func() { sendPubsub("Heartbeat", "heartbeat", "720") })
	}
	log.WithFields(log.Fields{"Heartbeat": "720", "Enabled": Conf.HbTwelveHour}).Info("Heartbeat")
	if Conf.HbTwentyfourHour == true {
		c.AddFunc("@every 24h", func() { sendPubsub("Heartbeat", "heartbeat", "1440") })
	}
	log.WithFields(log.Fields{"Heartbeat": "1440", "Enabled": Conf.HbTwentyfourHour}).Info("Heartbeat")
	c.Start()
}
