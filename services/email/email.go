package main

import (
	"flag"
	"github.com/nsqio/go-nsq"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/yguilai/timetable-micro/services/email/internal/config"
	"github.com/yguilai/timetable-micro/services/email/internal/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configFile = flag.String("f", "etc/email.yaml", "the config file")

func main() {
	flag.Parse()

	var cfg config.Config
	conf.MustLoad(*configFile, &cfg)

	consumer := MustNewConsumer(cfg.Nsq.Topic, cfg.Nsq.Channel)
	consumer.AddHandler(server.NewEmailServer(cfg.Mail))

	if err := consumer.ConnectToNSQLookupd(cfg.Nsq.Addr); err != nil {
		panic(err)
	}

	defer consumer.Stop()

	sign := make(chan os.Signal)
	signal.Notify(sign, syscall.SIGTERM, syscall.SIGINT)
	<-sign
}

func MustNewConsumer(topic, channel string) *nsq.Consumer {
	c := nsq.NewConfig()
	c.MaxAttempts = 10
	c.MaxInFlight = 150
	c.LookupdPollInterval = 15 * time.Second

	consumer, err := nsq.NewConsumer(topic, channel, c)
	if err != nil {
		panic(err)
	}
	return consumer
}
