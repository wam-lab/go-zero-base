package logic

import (
	"context"

	"github/yguilai/timetable-micro/gateway/internal/svc"
	"github/yguilai/timetable-micro/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CaptchaVerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) CaptchaVerifyLogic {
	return CaptchaVerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaVerifyLogic) CaptchaVerify(req types.CaptchaVerifyReq) (*types.CaptchaVerifyResp, error) {
	// todo: add your logic here and delete this line

	return &types.CaptchaVerifyResp{}, nil
}
