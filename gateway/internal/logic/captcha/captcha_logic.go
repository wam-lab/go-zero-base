package logic

import (
	"context"

	"github/yguilai/timetable-micro/gateway/internal/svc"
	"github/yguilai/timetable-micro/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) CaptchaLogic {
	return CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaLogic) Captcha() (*types.CaptchaResp, error) {
	// todo: add your logic here and delete this line

	return &types.CaptchaResp{}, nil
}
