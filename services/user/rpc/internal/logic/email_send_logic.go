package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yguilai/timetable-micro/common"
	"math/rand"
	"strings"
	"time"

	"github.com/yguilai/timetable-micro/services/user/rpc/internal/svc"
	"github.com/yguilai/timetable-micro/services/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type EmailSendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailSendLogic {
	return &EmailSendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const EmailKeyPrefix = "email_verify_"

func (l *EmailSendLogic) EmailSend(in *user.EmailSendReq) (*user.EmailSendResp, error) {
	// generate random code
	code := GenerateVerifyCode(6)

	// cache to redis
	err := l.svcCtx.Redis.Setex(EmailKeyPrefix+in.Email, code, 60*3)
	if err != nil {
		return nil, err
	}

	// build send content
	ctt, err := json.Marshal(&common.EmailContent{
		Target: in.Email,
		Title:  "验证码",
		Body:   fmt.Sprintf("您的邮箱验证码是: <b>%s</b>", code),
	})
	if err != nil {
		return nil, err
	}

	// publish email to beanstalkd
	_, err = l.svcCtx.Producer.Delay(ctt, time.Millisecond*100)
	if err != nil {
		return nil, err
	}
	return &user.EmailSendResp{Ok: true}, nil
}

func GenerateVerifyCode(l int) string {
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < l; i++ {
		key := rand.Intn(2)
		switch key {
		case 0:
			sb.WriteString(fmt.Sprintf("%d", rand.Intn(10)))
		case 1:
			sb.WriteString(fmt.Sprintf("%c", rand.Intn(26)+65))
		}
	}
	return sb.String()
}
