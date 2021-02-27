package logic

import (
	"context"
	"github.com/yguilai/timetable-micro/services/user/rpc/userclient"

	"github.com/yguilai/timetable-micro/gateway/internal/svc"
	"github.com/yguilai/timetable-micro/gateway/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (*types.RegisterResp, error) {
	resp, err := l.svcCtx.UserRpc.Register(l.ctx, &userclient.RegisterReq{
		Email:    req.Email,
		Nickname: req.Nickname,
		Password: req.Password,
		Verify:   req.Verify,
	})

	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		BaseResp: types.NewOkResp(),
		Token:    convertLocalToken(resp.Token),
		User:     convertLocalUserModel(resp.User),
	}, nil
}
