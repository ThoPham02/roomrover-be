package svc

import (
	"roomrover/service/contract/function"
	"roomrover/service/inventory/api/internal/config"
	"roomrover/service/inventory/model"
	"roomrover/storage"
	"roomrover/sync"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config  config.Config
	ObjSync *sync.ObjSync
	Storage *storage.StorageClient

	HouseModel   model.HouseTblModel
	RoomModel    model.RoomTblModel
	AlbumModel   model.AlbumTblModel
	ServiceModel model.ServiceTblModel

	ContractFunction function.ContractFunction
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		ObjSync: sync.NewObjSync(0),
		Storage: storage.NewStorageClient(c.Storage.Endpoint, c.Storage.AccessKey, c.Storage.SecretKey),

		HouseModel:   model.NewHouseTblModel(sqlx.NewMysql(c.DataSource)),
		RoomModel:    model.NewRoomTblModel(sqlx.NewMysql(c.DataSource)),
		AlbumModel:   model.NewAlbumTblModel(sqlx.NewMysql(c.DataSource)),
		ServiceModel: model.NewServiceTblModel(sqlx.NewMysql(c.DataSource)),
	}
}

func (sc *ServiceContext) SetContractFunction(cf function.ContractFunction) {
	sc.ContractFunction = cf
}
