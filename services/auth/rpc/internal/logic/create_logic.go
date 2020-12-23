package logic

import (
	"context"

	"github/yguilai/timetable-micro/services/auth/rpc/auth"
	"github/yguilai/timetable-micro/services/auth/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *auth.JwtCreateReq) (*auth.JwtCreateResp, error) {
	// todo: add your logic here and delete this line

	return &auth.JwtCreateResp{}, nil
}
