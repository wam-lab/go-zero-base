package handler

import (
	"github.com/yguilai/timetable-micro/common/xgo"
	"github.com/yguilai/timetable-micro/jobs/email/internal/svc"
)

func JobRun(svc *svc.ServiceContext)  {
	xgo.Go(func() {
		emailSendMessageHandler(svc)
	})
}