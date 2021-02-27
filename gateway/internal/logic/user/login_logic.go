package logic

import (
	"context"
	"github.com/yguilai/timetable-micro/services/user/rpc/user"
	"github.com/yguilai/timetable-micro/services/user/rpc/userclient"

	"github.com/yguilai/timetable-micro/gateway/internal/svc"
	"github.com/yguilai/timetable-micro/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginResp, error) {
	resp, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginReq{
		Email:    req.Email,
		Type:     int32(req.Type),
		Password: req.Password,
		Key:      req.Key,
		Code:     req.Code,
	})
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		BaseResp: types.NewOkResp(),
		Token:    convertLocalToken(resp.Token),
		User:     convertLocalUserModel(resp.User),
	}, nil
}

func convertLocalToken(t *user.Token) types.Token {
	return types.Token{
		AccessToken:  t.AccessToken,
		AccessExpire: t.AccessExpire,
		RefreshAfter: t.RefreshAfter,
	}
}

func convertLocalUserModel(u *user.UserModel) types.UserBaseResp {
	return types.UserBaseResp{
		Id:       u.Id,
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		OwnWords: u.OwnWords,
	}
}
