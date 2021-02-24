package handler

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/yguilai/timetable-micro/jobs/email/internal/logic"
	"github.com/yguilai/timetable-micro/jobs/email/internal/svc"
)

func emailSendMessageHandler(ctx *svc.ServiceContext)  {
	rCtx := context.Background()
	l := logic.NewEmailSendMessageLogic(context.Background(), ctx)

	err := l.EmailSendMessage()

	if err != nil {
		logx.WithContext(rCtx).Errorf("[JobErr] %+v", err)
	}
}
