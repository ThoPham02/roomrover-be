package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UsersTblModel = (*customUsersTblModel)(nil)

type (
	// UsersTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersTblModel.
	UsersTblModel interface {
		usersTblModel
		withSession(session sqlx.Session) UsersTblModel
	}

	customUsersTblModel struct {
		*defaultUsersTblModel
	}
)

// NewUsersTblModel returns a model for the database table.
func NewUsersTblModel(conn sqlx.SqlConn) UsersTblModel {
	return &customUsersTblModel{
		defaultUsersTblModel: newUsersTblModel(conn),
	}
}

func (m *customUsersTblModel) withSession(session sqlx.Session) UsersTblModel {
	return NewUsersTblModel(sqlx.NewSqlConnFromSession(session))
}
