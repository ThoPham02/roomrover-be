package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"roomrover/service/contract/api/internal/logic"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
)

func FilterContractHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FilterContractReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFilterContractLogic(r.Context(), svcCtx)
		resp, err := l.FilterContract(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
