package logic

import (
	"context"

	"github/yguilai/timetable-micro/services/jwt/rpc/internal/svc"
	"github/yguilai/timetable-micro/services/jwt/rpc/jwt"

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
	// todo: add your logic here and delete this line

	return &jwt.JwtVerifyResp{}, nil
}
