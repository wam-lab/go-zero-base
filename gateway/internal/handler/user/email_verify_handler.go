package handler

import (
	"net/http"

	"github/yguilai/timetable-micro/gateway/internal/logic/user"
	"github/yguilai/timetable-micro/gateway/internal/svc"
	"github/yguilai/timetable-micro/gateway/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func EmailVerifyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailVerifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEmailVerifyLogic(r.Context(), ctx)
		resp, err := l.EmailVerify(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
