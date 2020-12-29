// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

//go:generate mockgen -destination ./user_mock.go -package userclient -source $GOFILE

package userclient

import (
	"context"

	"github.com/yguilai/timetable-micro/services/user/rpc/user"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	PingReq        = user.PingReq
	Token          = user.Token
	RegisterReq    = user.RegisterReq
	RegisterResp   = user.RegisterResp
	LoginReq       = user.LoginReq
	LoginResp      = user.LoginResp
	PongResp       = user.PongResp
	UserModel      = user.UserModel
	EmailSendReq   = user.EmailSendReq
	EmailSendResp  = user.EmailSendResp
	EmailExistReq  = user.EmailExistReq
	EmailExistResp = user.EmailExistResp

	User interface {
		Ping(ctx context.Context, in *PingReq) (*PongResp, error)
		Register(ctx context.Context, in *RegisterReq) (*RegisterResp, error)
		Login(ctx context.Context, in *LoginReq) (*LoginResp, error)
		EmailSend(ctx context.Context, in *EmailSendReq) (*EmailSendResp, error)
		EmailExist(ctx context.Context, in *EmailExistReq) (*EmailExistResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Ping(ctx context.Context, in *PingReq) (*PongResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Ping(ctx, in)
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterReq) (*RegisterResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq) (*LoginResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in)
}

func (m *defaultUser) EmailSend(ctx context.Context, in *EmailSendReq) (*EmailSendResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.EmailSend(ctx, in)
}

func (m *defaultUser) EmailExist(ctx context.Context, in *EmailExistReq) (*EmailExistResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.EmailExist(ctx, in)
}
