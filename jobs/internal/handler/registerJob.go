package handler

import (
	"context"

	"github.com/tal-tech/go-zero/core/service"
	"github.com/yguilai/timetable-micro/jobs/email/internal/logic"
	"github.com/yguilai/timetable-micro/jobs/email/internal/svc"
)

// RegisterJob register all jobs
func RegisterJob(svc *svc.ServiceContext, group *service.ServiceGroup) {
	// register email job
	group.Add(logic.NewEmailSendMessageLogic(context.Background(), svc))

	group.Start()
}
