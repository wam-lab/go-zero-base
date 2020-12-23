package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"github/yguilai/timetable-micro/gateway/internal/config"
	"github/yguilai/timetable-micro/services/auth/rpc/authclient"
	"github/yguilai/timetable-micro/services/captcha/rpc/captchaclient"
	"github/yguilai/timetable-micro/services/user/rpc/userclient"
)

type ServiceContext struct {
	Config     config.Config
	CaptchaRpc captchaclient.Captcha
	UserRpc    userclient.User
	AuthRpc    authclient.Auth
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CaptchaRpc: captchaclient.NewCaptcha(zrpc.MustNewClient(c.CaptchaRpc)),
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		AuthRpc:    authclient.NewAuth(zrpc.MustNewClient(c.AuthRpc)),
	}
}
