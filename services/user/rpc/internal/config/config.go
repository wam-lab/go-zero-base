package config

import (
	"github.com/tal-tech/go-queue/dq"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	RedisConf  redis.RedisConf
	Nsq        struct {
		Addr  string
		Topic string
	}
	JwtRpc     zrpc.RpcClientConf
	CaptchaRpc zrpc.RpcClientConf
	Beanstalks []dq.Beanstalk
}
