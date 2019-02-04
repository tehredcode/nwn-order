package rds

import (
	"os"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// SendPubsub func
func SendPubsub(LogMessage string, PubsubChannel string, PubsubMessage string) {
	r := redis.NewClient(&redis.Options{Addr: os.Getenv("NWN_ORDER_REDIS_HOST") + ":" + os.Getenv("NWN_ORDER_REDIS_PORT")})
	err := r.Publish(PubsubChannel, PubsubMessage).Err()
	if err != nil {
		log.WithFields(log.Fields{"Channel": PubsubChannel, "Message": PubsubMessage}).Error(LogMessage)
		panic(err)
	}
	log.WithFields(log.Fields{"Channel": PubsubChannel, "Message": PubsubMessage}).Info(LogMessage)
}
