package handler

import (
	"net/http"

	"CloudDisk/core/internal/logic"
	"CloudDisk/core/internal/svc"
	"CloudDisk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MailCodeSendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MailCodeSendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMailCodeSendLogic(r.Context(), svcCtx)
		resp, err := l.MailCodeSend(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
