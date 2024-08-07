package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ HouseTblModel = (*customHouseTblModel)(nil)

type (
	// HouseTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHouseTblModel.
	HouseTblModel interface {
		houseTblModel
		withSession(session sqlx.Session) HouseTblModel
	}

	customHouseTblModel struct {
		*defaultHouseTblModel
	}
)

// NewHouseTblModel returns a model for the database table.
func NewHouseTblModel(conn sqlx.SqlConn) HouseTblModel {
	return &customHouseTblModel{
		defaultHouseTblModel: newHouseTblModel(conn),
	}
}

func (m *customHouseTblModel) withSession(session sqlx.Session) HouseTblModel {
	return NewHouseTblModel(sqlx.NewSqlConnFromSession(session))
}
