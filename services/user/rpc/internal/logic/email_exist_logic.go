package logic

import (
	"context"
	"github.com/yguilai/timetable-micro/services/user/model"

	"github.com/yguilai/timetable-micro/services/user/rpc/internal/svc"
	"github.com/yguilai/timetable-micro/services/user/rpc/user"

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
	var ec int64
	l.svcCtx.Db.Model(new(model.User)).Where("email = ?", in.Email).Count(&ec)
	if ec != 0 {
		return &user.EmailExistResp{Exist: true}, nil
	}

	return &user.EmailExistResp{Exist: false}, nil
}
