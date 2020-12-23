package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"github/yguilai/timetable-micro/gateway/internal/logic/captcha"
	"github/yguilai/timetable-micro/gateway/internal/svc"
)

func CaptchaHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewCaptchaLogic(r.Context(), ctx)
		resp, err := l.Captcha()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
