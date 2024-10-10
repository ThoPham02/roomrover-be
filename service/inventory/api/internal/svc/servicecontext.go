package svc

import (
	accountFunction "roomrover/service/account/function"
	contractFunction "roomrover/service/contract/function"
	"roomrover/service/inventory/api/internal/config"
	"roomrover/service/inventory/api/internal/middleware"
	"roomrover/service/inventory/model"
	"roomrover/storage"
	"roomrover/sync"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	UserTokenMiddleware rest.Middleware
	Config              config.Config
	ObjSync             *sync.ObjSync
	CldClient           *storage.CloudinaryClient

	HouseModel   model.HouseTblModel
	RoomModel    model.RoomTblModel
	AlbumModel   model.AlbumTblModel
	ServiceModel model.ServiceTblModel

	ContractFunction contractFunction.ContractFunction
	AccountFunction  accountFunction.AccountFunction
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		UserTokenMiddleware: middleware.NewUserTokenMiddleware().Handle,
		Config:              c,
		ObjSync:             sync.NewObjSync(0),
		CldClient:           storage.NewCloudinaryClient(c.Storage.CloudName, c.Storage.APIKey, c.Storage.APISecret, "inventory"),

		HouseModel:   model.NewHouseTblModel(sqlx.NewMysql(c.DataSource)),
		RoomModel:    model.NewRoomTblModel(sqlx.NewMysql(c.DataSource)),
		AlbumModel:   model.NewAlbumTblModel(sqlx.NewMysql(c.DataSource)),
		ServiceModel: model.NewServiceTblModel(sqlx.NewMysql(c.DataSource)),
	}
}

func (sc *ServiceContext) SetContractFunction(cf contractFunction.ContractFunction) {
	sc.ContractFunction = cf
}
