package logic

import (
	"context"
	"github/yguilai/timetable-micro/common/errory"
	"github/yguilai/timetable-micro/services/captcha/rpc/captchaclient"

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
	verify, err := l.svcCtx.CaptchaRpc.CaptchaVerify(l.ctx, &captchaclient.CaptchaVerifyReq{
		Key:  req.Key,
		Code: req.Code,
	})
	if err != nil {
		return nil, err
	}

	if !verify.Ok {
		return nil, errory.ErrCaptchaVerifyFailed
	}

	return &types.CaptchaVerifyResp{
		BaseResp: types.NewOkResp(),
	}, nil
}
