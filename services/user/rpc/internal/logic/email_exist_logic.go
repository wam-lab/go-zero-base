package logic

import (
	"context"

	"github/yguilai/timetable-micro/services/user/rpc/internal/svc"
	"github/yguilai/timetable-micro/services/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type EmailExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailExistLogic {
	return &EmailExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmailExistLogic) EmailExist(in *user.EmailExistReq) (*user.EmailExistResp, error) {
	// todo: add your logic here and delete this line

	return &user.EmailExistResp{}, nil
}
