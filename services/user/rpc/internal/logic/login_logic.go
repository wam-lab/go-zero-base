package logic

import (
	"context"
	"github.com/yguilai/timetable-micro/common"
	"github.com/yguilai/timetable-micro/common/errory"
	"github.com/yguilai/timetable-micro/services/user/model"
	"github.com/yguilai/timetable-micro/services/user/rpc/internal/svc"
	"github.com/yguilai/timetable-micro/services/user/rpc/user"
	"gorm.io/gorm"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const (
	LoginWithPwd = 1 + iota
	LoginWithCode
)

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	var ua model.UserAuth
	if err := l.svcCtx.Db.Where("email = ?", in.Email).First(&ua).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errory.ErrUserNotFound
		}
		return nil, err
	}

	switch in.Type {
	case LoginWithCode:
		verifyCode, err := l.svcCtx.Redis.Get(EmailKeyPrefix + in.Email)
		if err != nil {
			return nil, err
		}
		// incorrect verification code
		if verifyCode != in.Password {
			return nil, errory.ErrUserCodeIncorrect
		}
	case LoginWithPwd:
		fallthrough
	default:
		// if type is not exist, as login with password
		// incorrect password
		if !common.CheckPassword(ua.Password, in.Password) {
			return nil, errory.ErrUserPwdIncorrect
		}
	}
	// gen jwt token
	t, err := generateJwtToken(l.ctx, l.svcCtx.JwtRpc, ua.UserID)
	if err != nil {
		return nil, err
	}
	return &user.LoginResp{Token: t}, nil
}
