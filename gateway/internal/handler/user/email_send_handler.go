package handler

import (
	"net/http"

	"github.com/yguilai/timetable-micro/gateway/internal/logic/user"
	"github.com/yguilai/timetable-micro/gateway/internal/svc"
	"github.com/yguilai/timetable-micro/gateway/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func EmailSendHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailSendReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEmailSendLogic(r.Context(), ctx)
		resp, err := l.EmailSend(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
