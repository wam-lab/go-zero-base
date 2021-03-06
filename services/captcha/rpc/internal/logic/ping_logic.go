package logic

import (
	"context"

	"github.com/yguilai/timetable-micro/services/captcha/rpc/captcha"
	"github.com/yguilai/timetable-micro/services/captcha/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *captcha.PingReq) (*captcha.PongResp, error) {
	return &captcha.PongResp{Pong: "pong: " + in.Ping}, nil
}
