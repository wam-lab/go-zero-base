package logic

import (
	"context"
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/yguilai/timetable-micro/common"
	"github.com/yguilai/timetable-micro/services/jwt/rpc/internal/svc"
	"github.com/yguilai/timetable-micro/services/jwt/rpc/jwt"

	"github.com/tal-tech/go-zero/core/logx"
)

type VerifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLogic {
	return &VerifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyLogic) Verify(in *jwt.JwtVerifyReq) (*jwt.JwtVerifyResp, error) {
	token, err := common.ParseToken(in.Token, l.svcCtx.Config.JwtAuth.AccessSecret)
	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwtGo.MapClaims); ok && token.Valid {
		return &jwt.JwtVerifyResp{Valid: true}, nil
	}
	return &jwt.JwtVerifyResp{Valid: false}, nil
}
