// Code generated by goctl. DO NOT EDIT!
// Source: jwt.proto

package server

import (
	"context"

	"github.com/yguilai/timetable-micro/services/jwt/rpc/internal/logic"
	"github.com/yguilai/timetable-micro/services/jwt/rpc/internal/svc"
	"github.com/yguilai/timetable-micro/services/jwt/rpc/jwt"
)

type JwtServer struct {
	svcCtx *svc.ServiceContext
}

func NewJwtServer(svcCtx *svc.ServiceContext) *JwtServer {
	return &JwtServer{
		svcCtx: svcCtx,
	}
}

func (s *JwtServer) Ping(ctx context.Context, in *jwt.Request) (*jwt.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *JwtServer) Create(ctx context.Context, in *jwt.JwtCreateReq) (*jwt.JwtCreateResp, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *JwtServer) Verify(ctx context.Context, in *jwt.JwtVerifyReq) (*jwt.JwtVerifyResp, error) {
	l := logic.NewVerifyLogic(ctx, s.svcCtx)
	return l.Verify(in)
}

func (s *JwtServer) Refresh(ctx context.Context, in *jwt.JwtRefreshReq) (*jwt.JwtRefreshResp, error) {
	l := logic.NewRefreshLogic(ctx, s.svcCtx)
	return l.Refresh(in)
}
