package logic

import (
	"context"
	"github.com/yguilai/timetable-micro/common/errory"
	"gorm.io/gorm"

	"github.com/yguilai/timetable-micro/services/user/rpc/internal/svc"
	"github.com/yguilai/timetable-micro/services/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *user.InfoReq) (*user.InfoResp, error) {
	var u user.UserModel
	db := l.svcCtx.Db
	if err := db.Where("id = ?", in.Id).First(&u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errory.ErrUserNotFound
		}
		return nil, err
	}

	return &user.InfoResp{
		User: &u,
	}, nil
}
