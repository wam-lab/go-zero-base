package logic

import (
	"context"
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"

	"github/yguilai/timetable-micro/services/jwt/rpc/internal/svc"
	"github/yguilai/timetable-micro/services/jwt/rpc/jwt"

	"github.com/tal-tech/go-zero/core/logx"
)

type RefreshLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshLogic) Refresh(in *jwt.JwtRefreshReq) (*jwt.JwtRefreshResp, error) {
	ac := l.svcCtx.Config.Auth
	token, err := parseToken(in.Token, ac.AccessSecret)
	if err != nil {
		return nil, err
	}

	// token claims assert to MapClaims
	// reset exp and iat field of MapClaims
	if claims, ok := token.Claims.(jwtGo.MapClaims); ok && token.Valid {
		now := time.Now().Unix()
		claims["exp"] = now + ac.AccessExpire
		claims["iat"] = now

		token.Claims = claims
		// resign token
		newTokenString, err := token.SignedString([]byte(ac.AccessSecret))
		if err != nil {
			return nil, err
		}
		return &jwt.JwtRefreshResp{Token: &jwt.Token{
			AccessToken:  newTokenString,
			AccessExpire: now + ac.AccessExpire,
			RefreshAfter: now + ac.AccessExpire/2,
		}}, nil
	}

	return nil, errors.New("Refresh token failed")
}
