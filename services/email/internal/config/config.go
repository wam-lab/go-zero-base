package config

import "github.com/yguilai/timetable-micro/services/email/internal/server"

type Config struct {
	Name string
	Nsq  struct {
		Addr    string
		Topic   string
		Channel string
	}
	Mail server.Config
}
