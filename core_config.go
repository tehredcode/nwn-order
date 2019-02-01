package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config struct
type config struct {
	// core
	OrderPort string `env:"NWN_ORDER_PORT" envDefault:"5750"`
	// redis
	RedisPort string `env:"NWN_ORDER_REDIS_PORT" envDefault:"6379"`
	// module name
	ModuleName string `env:"NWN_ORDER_MODULE_NAME" envDefault:"DockerDemo"`
	// Webhooks
	GithubWebhookSecret    string `env:"NWN_ORDER_GITHUB_WEBHOOK_SECRET" envDefault:""`
	DockerhubWebhookSecret string `env:"NWN_ORDER_DOCKERHUB_WEBHOOK_SECRET" envDefault:""`

	// Discord
	PluginDiscord  bool   `env:"NWN_ORDER_PLUGIN_DISCORD_ENABLED" envDefault:"1"`
	DiscordBotKey  string `env:"NWN_ORDER_PLUGIN_DISCORD_BOT_KEY" envDefault:""`
	DiscordBotRoom string `env:"NWN_ORDER_PLUGIN_DISCOD_BOT_ROOM" envDefault:""`
	// Heartbeat
	PluginHearbeat   bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_ENABLED" envDefault:"1"`
	HbVerbose        bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_VERBOSE" envDefault:"false"`
	HbOneMinute      bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_ONE_MINUTE" envDefault:"true"`
	HbFiveMinute     bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_FIVE_MINUTE" envDefault:"true"`
	HbThirtyMinute   bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_THIRTY_MINUTE" envDefault:"true"`
	HbOneHour        bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_ONE_HOUR" envDefault:"true"`
	HbSixHour        bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_SIX_HOUR" envDefault:"true"`
	HbTwelveHour     bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_TWELVE_HOUR" envDefault:"true"`
	HbTwentyfourHour bool `env:"NWN_ORDER_PLUGIN_HEARTBEAT_TWENTYFOUR_HOUR" envDefault:"true"`
	// Logs
	PluginLogs   bool   `env:"NWN_ORDER_PLUGIN_LOG_ENABLED" envDefault:"1"`
	Pluginegress string `env:"NWN_ORDER_PLUGIN_LOG_EGRESS" envDefault:""`
}

// Create a new config instance.
var (
	Conf *config
)

// Read the config file from the current directory and marshal
// into the conf config struct.
func getConf() *config {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &config{}
	err = viper.Unmarshal(Conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}
