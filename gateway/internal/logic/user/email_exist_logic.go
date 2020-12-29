package logic

import (
	"context"

	"github.com/yguilai/timetable-micro/gateway/internal/svc"
	"github.com/yguilai/timetable-micro/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EmailExistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) EmailExistLogic {
	return EmailExistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailExistLogic) EmailExist(req types.EmailExistReq) (*types.EmailExistResp, error) {
	// todo: add your logic here and delete this line

	return &types.EmailExistResp{}, nil
}
