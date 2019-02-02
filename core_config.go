package main

import (
	"fmt"
	"github.com/caarlos0/env" 
)

// Config struct
type Config struct {
	OrderPort string `env:"NWN_ORDER_PORT" envDefault:"5750"`
	RedisPort string `env:"NWN_ORDER_REDIS_PORT" envDefault:"6379"`
	RedisHost string `env:"NWN_ORDER_REDIS_HOST" envDefault:"redis"`
	ModuleName string `env:"NWN_ORDER_MODULE_NAME"`
	BitbucketWebhookSecret string `env:"NWN_ORDER_BITBUCKET_WEBHOOK_SECRET"`
	GithubWebhookSecret    string `env:"NWN_ORDER_GITHUB_WEBHOOK_SECRET"`
	DockerhubWebhookSecret string `env:"NWN_ORDER_DOCKERHUB_WEBHOOK_SECRET"`
	PluginDiscord  bool   `env:"NWN_ORDER_PLUGIN_DISCORD_ENABLED"`
	DiscordBotKey  string `env:"NWN_ORDER_PLUGIN_DISCORD_BOT_KEY"`
	DiscordBotRoom string `env:"NWN_ORDER_PLUGIN_DISCOD_BOT_ROOM"`
	PluginHearbeat   bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_ENABLED"`
	HbVerbose        bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_VERBOSE"`
	HbOneMinute      bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_ONE_MINUTE"`
	HbFiveMinute     bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_FIVE_MINUTE"`
	HbThirtyMinute   bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_THIRTY_MINUTE"`
	HbOneHour        bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_ONE_HOUR"`
	HbSixHour        bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_SIX_HOUR"`
	HbTwelveHour     bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_TWELVE_HOUR"`
	HbTwentyfourHour bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_TWENTYFOUR_HOUR"`
	PluginLogs   bool   `env:"NWN_ORDER_PLUGIN_LOG_ENABLED"`
	Pluginegress string `env:"NWN_ORDER_PLUGIN_LOG_EGRESS"`
}

func initConf(){
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
