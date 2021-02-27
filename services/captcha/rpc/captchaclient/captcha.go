// Code generated by goctl. DO NOT EDIT!
// Source: captcha.proto

//go:generate mockgen -destination ./captcha_mock.go -package captchaclient -source $GOFILE

package captchaclient

import (
	"context"

	"github.com/yguilai/timetable-micro/services/captcha/rpc/captcha"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	CaptchaResp       = captcha.CaptchaResp
	CaptchaVerifyReq  = captcha.CaptchaVerifyReq
	CaptchaVerifyResp = captcha.CaptchaVerifyResp
	PingReq           = captcha.PingReq
	PongResp          = captcha.PongResp
	CaptchaReq        = captcha.CaptchaReq

	Captcha interface {
		Ping(ctx context.Context, in *PingReq) (*PongResp, error)
		CaptchaOne(ctx context.Context, in *CaptchaReq) (*CaptchaResp, error)
		CaptchaVerify(ctx context.Context, in *CaptchaVerifyReq) (*CaptchaVerifyResp, error)
	}

	defaultCaptcha struct {
		cli zrpc.Client
	}
)

func NewCaptcha(cli zrpc.Client) Captcha {
	return &defaultCaptcha{
		cli: cli,
	}
}

func (m *defaultCaptcha) Ping(ctx context.Context, in *PingReq) (*PongResp, error) {
	client := captcha.NewCaptchaClient(m.cli.Conn())
	return client.Ping(ctx, in)
}

func (m *defaultCaptcha) CaptchaOne(ctx context.Context, in *CaptchaReq) (*CaptchaResp, error) {
	client := captcha.NewCaptchaClient(m.cli.Conn())
	return client.CaptchaOne(ctx, in)
}

func (m *defaultCaptcha) CaptchaVerify(ctx context.Context, in *CaptchaVerifyReq) (*CaptchaVerifyResp, error) {
	client := captcha.NewCaptchaClient(m.cli.Conn())
	return client.CaptchaVerify(ctx, in)
}
