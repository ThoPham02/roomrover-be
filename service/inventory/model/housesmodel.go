package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ HousesModel = (*customHousesModel)(nil)

type (
	// HousesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHousesModel.
	HousesModel interface {
		housesModel
		withSession(session sqlx.Session) HousesModel
	}

	customHousesModel struct {
		*defaultHousesModel
	}
)

// NewHousesModel returns a model for the database table.
func NewHousesModel(conn sqlx.SqlConn) HousesModel {
	return &customHousesModel{
		defaultHousesModel: newHousesModel(conn),
	}
}

func (m *customHousesModel) withSession(session sqlx.Session) HousesModel {
	return NewHousesModel(sqlx.NewSqlConnFromSession(session))
}
