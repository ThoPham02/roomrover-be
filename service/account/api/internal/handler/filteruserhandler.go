package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"roomrover/service/account/api/internal/logic"
	"roomrover/service/account/api/internal/svc"
	"roomrover/service/account/api/internal/types"
)

func FilterUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FilterUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFilterUserLogic(r.Context(), svcCtx)
		resp, err := l.FilterUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
