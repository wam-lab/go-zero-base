package logic

import (
	"context"
	"github.com/yguilai/timetable-micro/common/cons"
	"github.com/yguilai/timetable-micro/services/user/rpc/userclient"

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
	resp, err := l.svcCtx.UserRpc.EmailExist(l.ctx, &userclient.EmailExistReq{Email: req.Email})
	if err != nil {
		return nil, err
	}

	br := types.NewBaseResp(cons.EmailNotExists, cons.EmailNotExistsMsg)
	if resp.Exist {
		br = types.NewBaseResp(cons.EmailExists, cons.EmailExistsMsg)
	}
	return &types.EmailExistResp{
		BaseResp: br,
	}, nil
}
