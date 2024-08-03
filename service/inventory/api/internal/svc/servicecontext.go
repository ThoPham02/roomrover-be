package svc

import (
	"roomrover/service/inventory/api/internal/config"
	"roomrover/service/inventory/model"
	"roomrover/sync"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	ObjSync *sync.ObjSync
	HouseModel model.HousesModel
	HouseServiceModel model.HouseServicesModel
	ClassModel model.ClassesModel
	ClassAlbum model.ClassAlbumsModel
	DistrictsModel model.DistrictsModel
	ProvincesModel model.ProvincesModel
	WardsModel model.WardsModel
	RoomsModel model.RoomsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		ObjSync: sync.NewObjSync(1),
		HouseModel: model.NewHousesModel(sqlx.NewMysql(c.DataSource)),
		HouseServiceModel: model.NewHouseServicesModel(sqlx.NewMysql(c.DataSource)),
		ClassModel: model.NewClassesModel(sqlx.NewMysql(c.DataSource)),
		ClassAlbum: model.NewClassAlbumsModel(sqlx.NewMysql(c.DataSource)),
		DistrictsModel: model.NewDistrictsModel(sqlx.NewMysql(c.DataSource)),
		ProvincesModel: model.NewProvincesModel(sqlx.NewMysql(c.DataSource)),
		WardsModel: model.NewWardsModel(sqlx.NewMysql(c.DataSource)),
		RoomsModel: model.NewRoomsModel(sqlx.NewMysql(c.DataSource)),
	}
}
