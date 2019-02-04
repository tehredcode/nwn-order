package config

import "os"

var (
	OrderPort              string
	RedisPort              string
	RedisHost              string
	ModuleName             string
	BitbucketWebhookSecret string
	GithubWebhookSecret    string
	DockerhubWebhookSecret string
	PluginDiscord          bool
	DiscordBotKey          string
	DiscordBotRoom         string
	PluginHearbeat         bool
	HbVerbose              bool
	HbOneMinute            bool
	HbFiveMinute           bool
	HbThirtyMinute         bool
	HbOneHour              bool
	HbSixHour              bool
	HbTwelveHour           bool
	HbTwentyfourHour       bool
	PluginLogs             bool
	Pluginegress           string
)

func initConf() {
	OrderPort = os.Getenv("NWN_ORDER_PORT")
	RedisPort = os.Getenv("NWN_ORDER_REDIS_PORT")
	RedisHost = os.Getenv("NWN_ORDER_REDIS_HOST")
	ModuleName = os.Getenv("NWN_ORDER_MODULE_NAME")
	BitbucketWebhookSecret = os.Getenv("NWN_ORDER_BITBUCKET_WEBHOOK_SECRET")
	GithubWebhookSecret = os.Getenv("NWN_ORDER_GITHUB_WEBHOOK_SECRET")
	DockerhubWebhookSecret = os.Getenv("NWN_ORDER_DOCKERHUB_WEBHOOK_SECRET")
	PluginDiscord = os.Getenv("NWN_ORDER_PLUGIN_DISCORD_ENABLED")
	DiscordBotKey = os.Getenv("NWN_ORDER_PLUGIN_DISCORD_BOT_KEY")
	DiscordBotRoom = os.Getenv("NWN_ORDER_PLUGIN_DISCOD_BOT_ROOM")
	PluginHearbeat = os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_ENABLED")
	HbVerbose = os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_VERBOSE")
	HbOneMinute = os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_ONE_MINUTE")
	HbFiveMinute = os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_FIVE_MINUTE")
	HbThirtyMinute = os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_THIRTY_MINUTE")
	HbOneHour = os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_ONE_HOUR")
	HbSixHour = os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_SIX_HOUR")
	HbTwelveHour = os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_TWELVE_HOUR")
	HbTwentyfourHour = os.Getenv("NWN_ORDER_PLUGIN_HEARTBEAT_TWENTYFOUR_HOUR")
	PluginLogs = os.Getenv("NWN_ORDER_PLUGIN_LOG_ENABLED")
	Pluginegress = os.Getenv("NWN_ORDER_PLUGIN_LOG_EGRESS")
}
