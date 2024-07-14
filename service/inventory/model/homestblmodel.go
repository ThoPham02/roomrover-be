package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ HomesTblModel = (*customHomesTblModel)(nil)

type (
	// HomesTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHomesTblModel.
	HomesTblModel interface {
		homesTblModel
		withSession(session sqlx.Session) HomesTblModel
	}

	customHomesTblModel struct {
		*defaultHomesTblModel
	}
)

// NewHomesTblModel returns a model for the database table.
func NewHomesTblModel(conn sqlx.SqlConn) HomesTblModel {
	return &customHomesTblModel{
		defaultHomesTblModel: newHomesTblModel(conn),
	}
}

func (m *customHomesTblModel) withSession(session sqlx.Session) HomesTblModel {
	return NewHomesTblModel(sqlx.NewSqlConnFromSession(session))
}
