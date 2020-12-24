package logic

import (
	"context"

	"github/yguilai/timetable-micro/services/jwt/rpc/internal/svc"
	"github/yguilai/timetable-micro/services/jwt/rpc/jwt"

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

func (l *RefreshLogic) Refresh(in *jwt.JwtRefreshReq) (*jwt.JwtRefreshResp, error) {
	// todo: add your logic here and delete this line

	return &jwt.JwtRefreshResp{}, nil
}
