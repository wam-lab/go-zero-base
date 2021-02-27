package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"github.com/yguilai/timetable-micro/gateway/internal/logic/user"
	"github.com/yguilai/timetable-micro/gateway/internal/svc"
)

func UserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), ctx)
		token := r.Header.Get("Authorization")
		resp, err := l.UserInfo(token)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
