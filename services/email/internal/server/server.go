package server

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"github.com/yguilai/timetable-micro/common"
	"gopkg.in/gomail.v2"
)

type EmailServer struct {
	c      Config
	dialer *gomail.Dialer
}

type Config struct {
	User string
	Pass string
	Host string
	Port int
}

func NewEmailServer(c Config) *EmailServer {
	return &EmailServer{
		c:      c,
		dialer: gomail.NewDialer(c.Host, c.Port, c.User, c.Pass),
	}
}

func (s *EmailServer) HandleMessage(msg *nsq.Message) error {
	var ec common.EmailContent
	err := json.Unmarshal(msg.Body, &ec)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", s.c.User)
	m.SetHeader("To", ec.Target)
	m.SetHeader("Subject", ec.Title)
	m.SetBody("text/html", ec.Body)
	
	if err := s.dialer.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
