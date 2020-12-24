package logic

import (
	"context"
	"github/yguilai/timetable-micro/services/captcha/rpc/captchaclient"

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
	captchaResp, err := l.svcCtx.CaptchaRpc.CaptchaOne(l.ctx, &captchaclient.CaptchaReq{})
	if err != nil {
		return nil, err
	}
	return &types.CaptchaResp{
		BaseResp: types.BaseResp{
			Code: 0,
			Msg:  "ok",
		},
		Key:  captchaResp.Key,
		Data: captchaResp.Data,
	}, nil
}
