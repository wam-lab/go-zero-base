package logic

import (
	"context"

	"github/yguilai/timetable-micro/services/captcha/rpc/captcha"
	"github/yguilai/timetable-micro/services/captcha/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CaptchaOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCaptchaOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaOneLogic {
	return &CaptchaOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CaptchaOneLogic) CaptchaOne(in *captcha.CaptchaReq) (*captcha.CaptchaResp, error) {
	// todo: add your logic here and delete this line

	return &captcha.CaptchaResp{}, nil
}
