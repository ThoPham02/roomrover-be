package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ProfilesTblModel = (*customProfilesTblModel)(nil)

type (
	// ProfilesTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProfilesTblModel.
	ProfilesTblModel interface {
		profilesTblModel
		withSession(session sqlx.Session) ProfilesTblModel
	}

	customProfilesTblModel struct {
		*defaultProfilesTblModel
	}
)

// NewProfilesTblModel returns a model for the database table.
func NewProfilesTblModel(conn sqlx.SqlConn) ProfilesTblModel {
	return &customProfilesTblModel{
		defaultProfilesTblModel: newProfilesTblModel(conn),
	}
}

func (m *customProfilesTblModel) withSession(session sqlx.Session) ProfilesTblModel {
	return NewProfilesTblModel(sqlx.NewSqlConnFromSession(session))
}
