package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	CaptchaRpc zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
	JwtRpc     zrpc.RpcClientConf
}
