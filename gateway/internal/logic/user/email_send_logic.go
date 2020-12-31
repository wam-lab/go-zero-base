package logic

import (
	"context"
	"github.com/yguilai/timetable-micro/common/cons"
	"github.com/yguilai/timetable-micro/services/user/rpc/userclient"

	"github.com/yguilai/timetable-micro/gateway/internal/svc"
	"github.com/yguilai/timetable-micro/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EmailSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) EmailSendLogic {
	return EmailSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailSendLogic) EmailSend(req types.EmailSendReq) (*types.EmailSendResp, error) {
	resp, err := l.svcCtx.UserRpc.EmailSend(l.ctx, &userclient.EmailSendReq{Email: req.Email})
	if err != nil {
		return nil, err
	}

	br := types.NewOkResp()
	if !resp.Ok {
		br = types.NewBaseResp(-1, cons.EmailSendingFailed)
	}
	return &types.EmailSendResp{
		BaseResp: br,
	}, nil
}
