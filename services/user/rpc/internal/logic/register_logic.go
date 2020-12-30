package logic

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/yguilai/timetable-micro/common"
	"github.com/yguilai/timetable-micro/common/errory"
	"github.com/yguilai/timetable-micro/services/jwt/rpc/jwt"
	"github.com/yguilai/timetable-micro/services/jwt/rpc/jwtclient"
	"github.com/yguilai/timetable-micro/services/user/model"
	"strconv"

	"github.com/yguilai/timetable-micro/services/user/rpc/internal/svc"
	"github.com/yguilai/timetable-micro/services/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// verify email code
	verifyInRedis, err := l.svcCtx.Redis.Get(EmailKeyPrefix + in.Email)
	if err != nil {
		return nil, err
	}
	if verifyInRedis != in.Verify {
		return nil, errory.ErrInvalidVerifyCode
	}

	// do register
	// create user model
	u := model.User{
		Email:    in.Email,
		Nickname: in.Nickname,
	}
	tx := l.svcCtx.Db.Begin()
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// crypt password
	password, err := common.GeneratePassword(in.Password)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// create user auth
	ua := model.UserAuth{
		UserID:   u.ID,
		Email:    u.Email,
		Password: password,
	}
	if err := tx.Create(&ua).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// generate jwt token
	// ignore create jwt error. So far, frontend should redirect to login page if token is nil
	token, _ := generateJwtToken(l.ctx, l.svcCtx.JwtRpc, u.ID)

	return &user.RegisterResp{
		User:  newUserModel(u),
		Token: token,
	}, nil
}

func newUserModel(u model.User) *user.UserModel {
	return &user.UserModel{
		Id:       int64(u.ID),
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		OwnWords: u.OwnWord,
		Gender:   int32(u.Gender),
		State:    int32(u.State),
		CreateAt: &timestamp.Timestamp{
			Seconds: u.CreatedAt.Unix(),
		},
		UpdateAt: &timestamp.Timestamp{
			Seconds: u.UpdatedAt.Unix(),
		},
	}
}

func newToken(t *jwt.Token) *user.Token {
	if t == nil {
		return nil
	}
	return &user.Token{
		AccessToken:  t.AccessToken,
		AccessExpire: t.AccessExpire,
		RefreshAfter: t.RefreshAfter,
	}
}

func generateJwtToken(ctx context.Context, rpc jwtclient.Jwt, uid uint) (*user.Token, error) {
	claims := make(map[string]string)
	claims["id"] = strconv.Itoa(int(uid))
	resp, err := rpc.Create(ctx, &jwt.JwtCreateReq{Claims: claims})
	if err != nil {
		return nil, err
	}
	return newToken(resp.Token), nil
}
