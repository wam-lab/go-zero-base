package config

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Redis   redis.RedisConf
	Captcha CaptchaConfig
}

type CaptchaConfig struct {
	Driver string `json:",default=math,options=math|string|number"`
	Length int    `json:",default=4"`
	Expire int    `json:",default=300"`
	Width  int    `json:",default=240"`
	Height int    `json:",default=80"`
}
