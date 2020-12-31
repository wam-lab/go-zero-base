package logic

import (
	"context"

	"github.com/yguilai/timetable-micro/services/user/rpc/internal/svc"
	"github.com/yguilai/timetable-micro/services/user/rpc/user"

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

func (l *PingLogic) Ping(in *user.PingReq) (*user.PongResp, error) {
	return &user.PongResp{Pong: "Pong: " + in.Ping}, nil
}
