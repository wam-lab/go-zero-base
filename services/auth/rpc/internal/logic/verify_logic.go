package logic

import (
	"context"

	"github/yguilai/timetable-micro/services/auth/rpc/auth"
	"github/yguilai/timetable-micro/services/auth/rpc/internal/svc"

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

func (l *VerifyLogic) Verify(in *auth.JwtVerifyReq) (*auth.JwtVerifyResp, error) {
	// todo: add your logic here and delete this line

	return &auth.JwtVerifyResp{}, nil
}
