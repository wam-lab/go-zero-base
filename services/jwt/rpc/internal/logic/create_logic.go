package logic

import (
	"context"
	"time"

	"github.com/yguilai/timetable-micro/services/jwt/rpc/internal/svc"
	"github.com/yguilai/timetable-micro/services/jwt/rpc/jwt"

	jwtGo "github.com/dgrijalva/jwt-go"
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
	ac := l.svcCtx.Config.JwtAuth

	now := time.Now().Unix()
	token, err := createJwtToken(now, ac.AccessExpire, ac.AccessSecret, in.Claims)
	if err != nil {
		return nil, err
	}
	return &jwt.JwtCreateResp{
		Token: &jwt.Token{
			AccessToken:  token,
			AccessExpire: now + ac.AccessExpire,
			RefreshAfter: now + ac.AccessExpire/2,
		},
	}, nil
}

func createJwtToken(iat, expire int64, secret string, payloads map[string]string) (string, error) {
	claims := make(jwtGo.MapClaims)
	claims["exp"] = iat + expire
	claims["iat"] = iat

	for k, v := range payloads {
		claims[k] = v
	}
	token := jwtGo.New(jwtGo.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}
