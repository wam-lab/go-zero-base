package handler

import (
	"context"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"github/yguilai/timetable-micro/gateway/internal/logic/jwt"
	"github/yguilai/timetable-micro/gateway/internal/svc"
)

func TokenRefreshHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c := context.WithValue(r.Context(), "token", r.Header.Get("Authorization"))
		l := logic.NewTokenRefreshLogic(c, ctx)
		resp, err := l.TokenRefresh()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
