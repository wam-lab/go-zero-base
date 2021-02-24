package svc

import (
	"github.com/tal-tech/go-queue/dq"
	"github.com/yguilai/timetable-micro/jobs/email/internal/config"
	"gopkg.in/gomail.v2"
)

type ServiceContext struct {
	Config   config.Config
	Consumer dq.Consumer
	Dialer   *gomail.Dialer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Consumer: dq.NewConsumer(c.DqConf),
		Dialer:   gomail.NewDialer(c.Mail.Host, c.Mail.Port, c.Mail.User, c.Mail.Pass),
	}
}
