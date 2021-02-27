package logic

import (
	"context"
	"github.com/yguilai/timetable-micro/common"
	"github.com/yguilai/timetable-micro/common/errory"
	"github.com/yguilai/timetable-micro/services/user/rpc/userclient"
	"strconv"

	"github.com/yguilai/timetable-micro/gateway/internal/svc"
	"github.com/yguilai/timetable-micro/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(token string) (*types.UserInfoResp, error) {
	claims, err := common.ParseTokenClaims(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return nil, err
	}

	var id int64
	if v, ok := claims["id"]; !ok {
		return nil, errory.ErrInvalidToken
	} else {
		id, err = strconv.ParseInt(v.(string), 10, 64)
	}

	resp, err := l.svcCtx.UserRpc.Info(l.ctx, &userclient.InfoReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResp{
		UserBaseResp: convertLocalUserModel(resp.User),
	}, nil
}
