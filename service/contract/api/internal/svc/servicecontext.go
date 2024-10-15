package svc

import (
	"roomrover/service/contract/api/internal/config"
	"roomrover/service/contract/api/internal/middleware"
	"roomrover/service/contract/model"
	"roomrover/sync"

	accountFunc "roomrover/service/account/function"
	inventFunc "roomrover/service/inventory/function"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config              config.Config
	UserTokenMiddleware rest.Middleware
	ObjSync             sync.ObjSync

	ContractModel      model.ContractTblModel
	ContractRedis      model.ContractRedis
	PaymentModel       model.PaymentTblModel
	PaymentDetailModel model.PaymentDetailTblModel
	PaymentRenterModel model.PaymentRenterTblModel
	BillModel          model.BillTblModel
	BillDetailModel    model.BillDetailTblModel
	BillPayModel       model.BillPayTblModel

	AccountFunction accountFunc.AccountFunction
	InventFunction  inventFunc.InventoryFunction
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.RedisCache.Host + ":" + c.RedisCache.Port,
		Password: c.RedisCache.Password,
		DB:       c.RedisCache.DB,
	})

	return &ServiceContext{
		Config:              c,
		UserTokenMiddleware: middleware.NewUserTokenMiddleware().Handle,
		ObjSync:             *sync.NewObjSync(0),
		ContractRedis:       model.NewContractRedisClient(rdb),
		ContractModel:       model.NewContractTblModel(sqlx.NewMysql(c.DataSource)),
		PaymentModel:        model.NewPaymentTblModel(sqlx.NewMysql(c.DataSource)),
		PaymentDetailModel:  model.NewPaymentDetailTblModel(sqlx.NewMysql(c.DataSource)),
		PaymentRenterModel:  model.NewPaymentRenterTblModel(sqlx.NewMysql(c.DataSource)),
		BillModel:           model.NewBillTblModel(sqlx.NewMysql(c.DataSource)),
		BillDetailModel:     model.NewBillDetailTblModel(sqlx.NewMysql(c.DataSource)),
		BillPayModel:        model.NewBillPayTblModel(sqlx.NewMysql(c.DataSource)),
	}
}

func (sc *ServiceContext) SetAccountFunction(accountFunction accountFunc.AccountFunction) {
	sc.AccountFunction = accountFunction
}

func (sc *ServiceContext) SetInventFunction(inventFunction inventFunc.InventoryFunction) {
	sc.InventFunction = inventFunction
}
