// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"roomrover/service/contract/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserTokenMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPut,
					Path:    "/:id/confirm",
					Handler: ConfirmContractHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: CreateContractHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/:id",
					Handler: UpdateContractHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: GetContractHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id",
					Handler: DeleteContractHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/filter",
					Handler: FilterContractHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/:id/status",
					Handler: UpdateContractStatusHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/contract"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserTokenMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/filter",
					Handler: FilterBillHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: GetBillDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/:id",
					Handler: UpdateBillDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/pay",
					Handler: CreateBillPayHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/pay/:id",
					Handler: DeleteBillPayHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/:id/status",
					Handler: UpdateBillStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/zalo",
					Handler: ZaloPaymentHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/bill"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/zalo/callback",
				Handler: ZaloPaymentCallbackHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/bill"),
	)
}
