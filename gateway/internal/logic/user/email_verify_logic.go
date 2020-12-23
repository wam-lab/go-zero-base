package logic

import (
	"context"

	"github/yguilai/timetable-micro/gateway/internal/svc"
	"github/yguilai/timetable-micro/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EmailVerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) EmailVerifyLogic {
	return EmailVerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailVerifyLogic) EmailVerify(req types.EmailVerifyReq) (*types.EmailVerifyResp, error) {
	// todo: add your logic here and delete this line

	return &types.EmailVerifyResp{}, nil
}
