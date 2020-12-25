package logic

import (
	"context"
	"github/yguilai/timetable-micro/common"
	"github/yguilai/timetable-micro/services/jwt/rpc/jwtclient"

	"github/yguilai/timetable-micro/gateway/internal/svc"
	"github/yguilai/timetable-micro/gateway/internal/types"

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
	if token, ok := l.ctx.Value("token").(string); ok && token != "" {
		refresh, err := l.svcCtx.JwtRpc.Refresh(l.ctx, &jwtclient.JwtRefreshReq{
			Token: token,
		})
		if err != nil {
			return nil, err
		}
		return &types.Token{
			AccessToken:  refresh.Token.AccessToken,
			AccessExpire: refresh.Token.AccessExpire,
			RefreshAfter: refresh.Token.RefreshAfter,
		}, nil
	}
	return nil, common.ErrInvalidToken
}
