package logic

import (
	"context"
	"encoding/json"

	"github.com/yguilai/timetable-micro/common"
	"gopkg.in/gomail.v2"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/threading"
	"github.com/yguilai/timetable-micro/jobs/email/internal/svc"
)

type EmailSendMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailSendMessageLogic {
	return &EmailSendMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailSendMessageLogic) Start() {
	logx.Info("Start consume\n")

	threading.GoSafe(func() {
		l.svcCtx.Consumer.Consume(func(body []byte) {
			var mail common.EmailContent
			err := json.Unmarshal(body, &mail)
			if err != nil {
				l.Logger.Error(err)
				return
			}

			m := gomail.NewMessage()
			m.SetHeader("From", l.svcCtx.Config.Mail.User)
			m.SetHeader("To", mail.Target)
			m.SetHeader("Subject", mail.Title)
			m.SetBody("text/html", mail.Body)

			if err := l.svcCtx.Dialer.DialAndSend(m); err != nil {
				l.Logger.Error(err)
				return
			}
		})
	})
}

func (l *EmailSendMessageLogic) Stop() {
	logx.Info("Stop cosume\n")
}
