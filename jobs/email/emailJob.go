package main

import (
	"flag"
	"fmt"

	"github.com/yguilai/timetable-micro/jobs/email/internal/config"
	"github.com/yguilai/timetable-micro/jobs/email/internal/handler"
	"github.com/yguilai/timetable-micro/jobs/email/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/emailJob.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.JobRun(ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
