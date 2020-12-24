package config

import "github.com/tal-tech/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
