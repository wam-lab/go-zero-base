package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &types.EmailSendResp{}, nil
}
