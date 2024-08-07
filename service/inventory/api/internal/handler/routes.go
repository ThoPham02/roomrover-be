// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"roomrover/service/inventory/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// Create house
				Method:  http.MethodPost,
				Path:    "/house",
				Handler: CreateHouseHandler(serverCtx),
			},
			{
				// Get house
				Method:  http.MethodGet,
				Path:    "/house",
				Handler: GetHouseHandler(serverCtx),
			},
			{
				// Update house
				Method:  http.MethodPut,
				Path:    "/house",
				Handler: UpdateHouseHandler(serverCtx),
			},
			{
				// Delete house
				Method:  http.MethodDelete,
				Path:    "/house",
				Handler: DeleteHouseHandler(serverCtx),
			},
			{
				// Filter house
				Method:  http.MethodPost,
				Path:    "/house/filter",
				Handler: FilterHouseHandler(serverCtx),
			},
			{
				// Create room
				Method:  http.MethodPost,
				Path:    "/room",
				Handler: CreateRoomHandler(serverCtx),
			},
			{
				// Get room
				Method:  http.MethodGet,
				Path:    "/room",
				Handler: GetRoomHandler(serverCtx),
			},
			{
				// Update room
				Method:  http.MethodPut,
				Path:    "/room",
				Handler: UpdateRoomHandler(serverCtx),
			},
			{
				// Delete room
				Method:  http.MethodDelete,
				Path:    "/room",
				Handler: DeleteRoomHandler(serverCtx),
			},
			{
				// Get room by house
				Method:  http.MethodGet,
				Path:    "/room/house",
				Handler: GetRoomByHouseHandler(serverCtx),
			},
			{
				// Upload file house
				Method:  http.MethodPost,
				Path:    "/upload/house",
				Handler: UploadFileHouseHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/invent"),
	)
}
