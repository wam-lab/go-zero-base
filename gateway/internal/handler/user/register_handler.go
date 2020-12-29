package handler

import (
	"net/http"

	"github.com/yguilai/timetable-micro/gateway/internal/logic/user"
	"github.com/yguilai/timetable-micro/gateway/internal/svc"
	"github.com/yguilai/timetable-micro/gateway/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func RegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), ctx)
		resp, err := l.Register(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
