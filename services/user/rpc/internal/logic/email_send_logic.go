package logic

import (
	"context"

	"github/yguilai/timetable-micro/services/user/rpc/internal/svc"
	"github/yguilai/timetable-micro/services/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type EmailSendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailSendLogic {
	return &EmailSendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmailSendLogic) EmailSend(in *user.EmailSendReq) (*user.EmailSendResp, error) {
	// todo: add your logic here and delete this line

	return &user.EmailSendResp{}, nil
}
