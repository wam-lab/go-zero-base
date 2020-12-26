package logic

import (
	"context"
	"github/yguilai/timetable-micro/services/captcha/rpc/captcha"
	"github/yguilai/timetable-micro/services/captcha/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CaptchaVerifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCaptchaVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaVerifyLogic {
	return &CaptchaVerifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CaptchaVerifyLogic) CaptchaVerify(in *captcha.CaptchaVerifyReq) (*captcha.CaptchaVerifyResp, error) {
	// clear: true, ignore if verify successfully
	verify := l.svcCtx.Store.Verify(in.Key, in.Code, true)
	return &captcha.CaptchaVerifyResp{Ok: verify}, nil
}
