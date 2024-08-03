package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProvincesModel = (*customProvincesModel)(nil)

type (
	// ProvincesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProvincesModel.
	ProvincesModel interface {
		provincesModel
		withSession(session sqlx.Session) ProvincesModel
	}

	customProvincesModel struct {
		*defaultProvincesModel
	}
)

// NewProvincesModel returns a model for the database table.
func NewProvincesModel(conn sqlx.SqlConn) ProvincesModel {
	return &customProvincesModel{
		defaultProvincesModel: newProvincesModel(conn),
	}
}

func (m *customProvincesModel) withSession(session sqlx.Session) ProvincesModel {
	return NewProvincesModel(sqlx.NewSqlConnFromSession(session))
}
