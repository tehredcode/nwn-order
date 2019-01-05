package main

type config struct {
	RedisPort        string `env:"NWN_ORDER_REDIS_PORT" envDefault:"6379"`
	OrderPort        string `env:"NWN_ORDER_PORT" envDefault:"5750"`
	HbVerbose        bool   `env:"NWN_ORDER_HB_VERBOSE" envDefault:"true"`
	HbOneMinute      bool   `env:"NWN_ORDER_HB_ONE_MINUTE" envDefault:"true"`
	HbFiveMinute     bool   `env:"NWN_ORDER_HB_FIVE_MINUTE" envDefault:"true"`
	HbThirtyMinute   bool   `env:"NWN_ORDER_HB_THIRTY_MINUTE" envDefault:"true"`
	HbOneHour        bool   `env:"NWN_ORDER_HB_ONE_HOUR" envDefault:"true"`
	HbSixHour        bool   `env:"NWN_ORDER_HB_SIX_HOUR" envDefault:"true"`
	HbTwelveHour     bool   `env:"NWN_ORDER_HB_TWELVE_HOUR" envDefault:"true"`
	HbTwentyfourHour bool   `env:"NWN_ORDER_HB_TWENTYFOUR_HOUR" envDefault:"true"`
}
