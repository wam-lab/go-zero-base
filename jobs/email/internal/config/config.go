package config

import (
	"github.com/tal-tech/go-queue/dq"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mail struct{
		User string
		Pass string
		Host string
		Port int
	}
	DqConf dq.DqConf
}
