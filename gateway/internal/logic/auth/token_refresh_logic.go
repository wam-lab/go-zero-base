package logic

import (
	"context"

	"github.com/yguilai/timetable-micro/gateway/internal/svc"
	"github.com/yguilai/timetable-micro/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type TokenRefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) TokenRefreshLogic {
	return TokenRefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenRefreshLogic) TokenRefresh() (*types.Token, error) {
	// todo: add your logic here and delete this line

	return &types.Token{}, nil
}
