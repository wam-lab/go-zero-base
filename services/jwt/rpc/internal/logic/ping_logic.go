package logic

import (
	"context"

	"github.com/yguilai/timetable-micro/services/jwt/rpc/internal/svc"
	"github.com/yguilai/timetable-micro/services/jwt/rpc/jwt"

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

func (l *PingLogic) Ping(in *jwt.Request) (*jwt.Response, error) {
	return &jwt.Response{Pong: "pong: " + in.Ping}, nil
}
