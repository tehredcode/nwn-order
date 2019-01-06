package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/robfig/cron"
	"github.com/caarlos0/env"
	"github.com/glendc/go-external-ip"
	"github.com/go-redis/redis"
	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
)

type config struct {
	RedisPort        string `env:"NWN_ORDER_REDIS_PORT" envDefault:"6379"`
	OrderPort        string `env:"NWN_ORDER_PORT" envDefault:"5750"`
	HbVerbose        bool   `env:"NWN_ORDER_HB_VERBOSE" envDefault:"false"`
	HbOneMinute      bool   `env:"NWN_ORDER_HB_ONE_MINUTE" envDefault:"true"`
	HbFiveMinute     bool   `env:"NWN_ORDER_HB_FIVE_MINUTE" envDefault:"true"`
	HbThirtyMinute   bool   `env:"NWN_ORDER_HB_THIRTY_MINUTE" envDefault:"true"`
	HbOneHour        bool   `env:"NWN_ORDER_HB_ONE_HOUR" envDefault:"true"`
	HbSixHour        bool   `env:"NWN_ORDER_HB_SIX_HOUR" envDefault:"true"`
	HbTwelveHour     bool   `env:"NWN_ORDER_HB_TWELVE_HOUR" envDefault:"true"`
	HbTwentyfourHour bool   `env:"NWN_ORDER_HB_TWENTYFOUR_HOUR" envDefault:"true"`
}

var (
	//RedisClient is a var
	RedisClient *redis.Client
)

func startPubsub() {
	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + cfg.RedisPort,
	})
	defer client.Close()

	pubSub := client.Subscribe(
		"heartbeat",
		"input",
		"debug",
		"github",
	)
	for {
		msg, _ := pubSub.ReceiveMessage()
		switch msg.Channel {
		case "heartbeat":

		case "input":
			go uuidGeneration()
		case "debug":
		}
	}
}

func uuidGeneration() {
	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + cfg.RedisPort,
	})
	defer client.Close()

	b := make([]byte, 16)
	_, err = rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x%x%x%x%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	if err := client.Set("system:uuid", uuid, 0).Err(); err != nil {
		panic(err)
	}

	pub := client.Publish("output", "uuid")
	if err = pub.Err(); err != nil {
		fmt.Print("PublishString() error", err)
	}
}

func sendPubsub(LogMessage string, PubsubChannel string, PubsubMessage string) {
	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: "redis:" + cfg.RedisPort,
	})
	defer client.Close()

	if err := client.Publish(PubsubChannel, PubsubMessage).Err(); err != nil {
		panic(err)
	}

	fmt.Println(LogMessage)
}

func heartbeatWebhook(ticker string, verbose bool) {
	t := time.Now()
	msg := ("O [" + t.Format("15:04:05") + "] [NWN_Order] Pubsub Event: channel=heartbeat message=" + ticker)
	sendPubsub(msg, "heartbeat", ticker)
}

func githubWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error validating request body: err=%s\n", err)
		return
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		fmt.Printf("could not parse webhook: err=%s\n", err)
		return
	}

	switch e := event.(type) {

	case *github.PushEvent:
		t := time.Now()
		msg := ("O [" + t.Format("15:04:05") + "] [NWN_Order] Webhook Event: channel=innwserver message=repoupdate | " + *e.Sender.Login + " made a commit to module repo")
		go sendPubsub(msg, "github", "commit")

	default:
		fmt.Printf("Only push events supported, unknown webhook event type %s\n", github.WebHookType(r))
		return
	}
}

func webpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func webserver() {
	cfg := config{}
	err := env.Parse(&cfg)

	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	consensus := externalip.DefaultConsensus(nil, nil)
	ip, _ := consensus.ExternalIP()
	http.HandleFunc("/webhook", githubWebhook)
	t := time.Now()
	fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: webserver started with external IP of " + ip.String() + ":" + cfg.OrderPort + ". webhooks need to be sent to /webhook")

	http.HandleFunc("/", webpage)
	log.Fatal(http.ListenAndServe(":"+cfg.OrderPort, nil))
}

func main() {
    done := make(chan bool)
	go initMain()
    <-done // Block forever
}

func initMain() {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	t := time.Now()
	log.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: Order has Started")

	conn, err := net.Dial("udp", "redis:6379")
	for retry := 1; err != nil; retry++ {
		trds := time.Now()
		s := strconv.Itoa(retry)
		fmt.Println("O [" + trds.Format("15:04:05") + "] [NWN_Order] Boot Event: Redis not connected | Retry attempt: " + s + " | 5 second sleep")
		if retry > 4 {
			fmt.Println("O [" + trds.Format("15:04:05") + "] [NWN_Order] Boot Event: Redis not connected | Exiting")
			os.Exit(1)
		}
		time.Sleep(5 * time.Second)
	}
	conn.Close()
	t = time.Now()
	fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: Redis connected")

	// start pubsub
	go startPubsub()
	fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: Pubsub started")

	// start webhook reciever
	go webserver()
	fmt.Println("O [" + t.Format("15:04:05") + "] [NWN_Order] Boot Event: Webserver started")

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

	for {
        fmt.Printf("%v+\n", time.Now())
        time.Sleep(time.Second)
    }
}