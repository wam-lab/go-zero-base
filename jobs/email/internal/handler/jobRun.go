package handler

import (
	"github.com/tal-tech/go-zero/core/threading"
	"github.com/yguilai/timetable-micro/jobs/email/internal/svc"
)

func JobRun(svc *svc.ServiceContext) {
	threading.GoSafe(func() {
		emailSendMessageHandler(svc)
	})
}
