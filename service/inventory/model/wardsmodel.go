package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ WardsModel = (*customWardsModel)(nil)

type (
	// WardsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWardsModel.
	WardsModel interface {
		wardsModel
		withSession(session sqlx.Session) WardsModel
	}

	customWardsModel struct {
		*defaultWardsModel
	}
)

// NewWardsModel returns a model for the database table.
func NewWardsModel(conn sqlx.SqlConn) WardsModel {
	return &customWardsModel{
		defaultWardsModel: newWardsModel(conn),
	}
}

func (m *customWardsModel) withSession(session sqlx.Session) WardsModel {
	return NewWardsModel(sqlx.NewSqlConnFromSession(session))
}
