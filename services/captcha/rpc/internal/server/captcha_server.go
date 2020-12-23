// Code generated by goctl. DO NOT EDIT!
// Source: captcha.proto

package server

import (
	"context"

	"github/yguilai/timetable-micro/services/captcha/rpc/captcha"
	"github/yguilai/timetable-micro/services/captcha/rpc/internal/logic"
	"github/yguilai/timetable-micro/services/captcha/rpc/internal/svc"
)

type CaptchaServer struct {
	svcCtx *svc.ServiceContext
}

func NewCaptchaServer(svcCtx *svc.ServiceContext) *CaptchaServer {
	return &CaptchaServer{
		svcCtx: svcCtx,
	}
}

func (s *CaptchaServer) CaptchaOne(ctx context.Context, in *captcha.CaptchaReq) (*captcha.CaptchaResp, error) {
	l := logic.NewCaptchaOneLogic(ctx, s.svcCtx)
	return l.CaptchaOne(in)
}

func (s *CaptchaServer) CaptchaVerify(ctx context.Context, in *captcha.CaptchaVerifyReq) (*captcha.CaptchaVerifyResp, error) {
	l := logic.NewCaptchaVerifyLogic(ctx, s.svcCtx)
	return l.CaptchaVerify(in)
}
