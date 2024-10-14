package svc

import (
	accountFunc "roomrover/service/account/function"
	contractFunc "roomrover/service/contract/function"
	"roomrover/service/inventory/api/internal/config"
	"roomrover/service/inventory/api/internal/middleware"
	"roomrover/service/inventory/model"
	"roomrover/storage"
	"roomrover/sync"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config              config.Config
	UserTokenMiddleware rest.Middleware
	ObjSync             *sync.ObjSync
	CldClient           *storage.CloudinaryClient

	RoomModel    model.RoomTblModel
	HouseModel   model.HouseTblModel
	AlbumModel   model.AlbumTblModel
	ServiceModel model.ServiceTblModel

	AccountFunction  accountFunc.AccountFunction
	ContractFunction contractFunc.ContractFunction
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		UserTokenMiddleware: middleware.NewUserTokenMiddleware().Handle,
		ObjSync:             sync.NewObjSync(1),
		CldClient:           storage.NewCloudinaryClient(c.Storage.CloudName, c.Storage.APIKey, c.Storage.APISecret, "inventory"),

		RoomModel:    model.NewRoomTblModel(sqlx.NewMysql(c.DataSource)),
		HouseModel:   model.NewHouseTblModel(sqlx.NewMysql(c.DataSource)),
		AlbumModel:   model.NewAlbumTblModel(sqlx.NewMysql(c.DataSource)),
		ServiceModel: model.NewServiceTblModel(sqlx.NewMysql(c.DataSource)),
	}
}

func (ctx *ServiceContext) SetContractFunction(contractFunc contractFunc.ContractFunction) {
	ctx.ContractFunction = contractFunc
}

func (ctx *ServiceContext) SetAccountFunction(accountFunc accountFunc.AccountFunction) {
	ctx.AccountFunction = accountFunc
}
