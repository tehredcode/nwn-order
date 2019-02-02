package main

import (
	"fmt"

	"github.com/caarlos0/env"
)

// Config struct
type Config struct {
	OrderPort              string `env:"NWN_ORDER_PORT" envDefault:"5750"`
	RedisPort              string `env:"NWN_ORDER_REDIS_PORT" envDefault:"6379"`
	RedisHost              string `env:"NWN_ORDER_REDIS_HOST" envDefault:"redis"`
	ModuleName             string `env:"NWN_ORDER_MODULE_NAME" envDefault:"redis"`
	BitbucketWebhookSecret string `env:"NWN_ORDER_BITBUCKET_WEBHOOK_SECRET" envDefault:"asd"`
	GithubWebhookSecret    string `env:"NWN_ORDER_GITHUB_WEBHOOK_SECRET" envDefault:"asd"`
	DockerhubWebhookSecret string `env:"NWN_ORDER_DOCKERHUB_WEBHOOK_SECRET" envDefault:"asd"`
	PluginDiscord          bool   `env:"NWN_ORDER_PLUGIN_DISCORD_ENABLED" envDefault:"0"`
	DiscordBotKey          string `env:"NWN_ORDER_PLUGIN_DISCORD_BOT_KEY" envDefault:"asd"`
	DiscordBotRoom         string `env:"NWN_ORDER_PLUGIN_DISCOD_BOT_ROOM" envDefault:"asd"`
	PluginHearbeat         bool   `env:"NWN_ORDER_PLUGIN_HEARTBEAT_ENABLED" envDefault:"1"`
	HbVerbose              bool   `env:"NWN_ORDER_PLUGIN_HEARTBEAT_VERBOSE" envDefault:"0"`
	HbOneMinute            bool   `env:"NWN_ORDER_PLUGIN_HEARTBEAT_ONE_MINUTE" envDefault:"1"`
	HbFiveMinute           bool   `env:"NWN_ORDER_PLUGIN_HEARTBEAT_FIVE_MINUTE" envDefault:"1"`
	HbThirtyMinute         bool   `env:"NWN_ORDER_PLUGIN_HEARTBEAT_THIRTY_MINUTE" envDefault:"1"`
	HbOneHour              bool   `env:"NWN_ORDER_PLUGIN_HEARTBEAT_ONE_HOUR" envDefault:"1"`
	HbSixHour              bool   `env:"NWN_ORDER_PLUGIN_HEARTBEAT_SIX_HOUR" envDefault:"1"`
	HbTwelveHour           bool   `env:"NWN_ORDER_PLUGIN_HEARTBEAT_TWELVE_HOUR" envDefault:"1"`
	HbTwentyfourHour       bool   `env:"NWN_ORDER_PLUGIN_HEARTBEAT_TWENTYFOUR_HOUR" envDefault:"1"`
	PluginLogs             bool   `env:"NWN_ORDER_PLUGIN_LOG_ENABLED" envDefault:"1"`
	Pluginegress           string `env:"NWN_ORDER_PLUGIN_LOG_EGRESS" envDefault:"idk"`
}

func initConf() {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
