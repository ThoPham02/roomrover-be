package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"roomrover/service/notification/api/internal/logic"
	"roomrover/service/notification/api/internal/svc"
	"roomrover/service/notification/api/internal/types"
)

func GetListNotificationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetListNotificationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetListNotificationLogic(r.Context(), svcCtx)
		resp, err := l.GetListNotification(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
