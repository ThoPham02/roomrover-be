package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"roomrover/service/contract/api/internal/logic"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
)

func ZaloPaymentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ZaloPaymentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewZaloPaymentLogic(r.Context(), svcCtx)
		resp, err := l.ZaloPayment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
