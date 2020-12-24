package logic

import (
	"context"

	"github/yguilai/timetable-micro/services/jwt/rpc/internal/svc"
	"github/yguilai/timetable-micro/services/jwt/rpc/jwt"

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

func (l *CreateLogic) Create(in *jwt.JwtCreateReq) (*jwt.JwtCreateResp, error) {
	// todo: add your logic here and delete this line

	return &jwt.JwtCreateResp{}, nil
}
