package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ HouseServicesModel = (*customHouseServicesModel)(nil)

type (
	// HouseServicesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHouseServicesModel.
	HouseServicesModel interface {
		houseServicesModel
		withSession(session sqlx.Session) HouseServicesModel
	}

	customHouseServicesModel struct {
		*defaultHouseServicesModel
	}
)

// NewHouseServicesModel returns a model for the database table.
func NewHouseServicesModel(conn sqlx.SqlConn) HouseServicesModel {
	return &customHouseServicesModel{
		defaultHouseServicesModel: newHouseServicesModel(conn),
	}
}

func (m *customHouseServicesModel) withSession(session sqlx.Session) HouseServicesModel {
	return NewHouseServicesModel(sqlx.NewSqlConnFromSession(session))
}
