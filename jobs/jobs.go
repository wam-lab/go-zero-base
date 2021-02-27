package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/yguilai/timetable-micro/jobs/email/internal/config"
	"github.com/yguilai/timetable-micro/jobs/email/internal/handler"
	"github.com/yguilai/timetable-micro/jobs/email/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/emailJob.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	group := service.NewServiceGroup()
	handler.RegisterJob(ctx, group)

	sign := make(chan os.Signal)
	signal.Notify(sign, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <- sign
		logx.Infof("Got a signal %s", s.String())
		switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:

		}
	}
}
