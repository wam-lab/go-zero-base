package logic

import (
	"context"

	"github/yguilai/timetable-micro/services/auth/rpc/auth"
	"github/yguilai/timetable-micro/services/auth/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type RefreshLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshLogic) Refresh(in *auth.JwtRefreshReq) (*auth.JwtRefreshResp, error) {
	// todo: add your logic here and delete this line

	return &auth.JwtRefreshResp{}, nil
}
