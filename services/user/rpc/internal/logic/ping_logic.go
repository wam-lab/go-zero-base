package logic

import (
	"context"

	"github/yguilai/timetable-micro/services/user/rpc/internal/svc"
	"github/yguilai/timetable-micro/services/user/rpc/user"

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
	// todo: add your logic here and delete this line

	return &user.PongResp{}, nil
}
