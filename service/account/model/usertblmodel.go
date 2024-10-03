package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserTblModel = (*customUserTblModel)(nil)

type (
	// UserTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserTblModel.
	UserTblModel interface {
		userTblModel
		withSession(session sqlx.Session) UserTblModel
	}

	customUserTblModel struct {
		*defaultUserTblModel
	}
)

// NewUserTblModel returns a model for the database table.
func NewUserTblModel(conn sqlx.SqlConn) UserTblModel {
	return &customUserTblModel{
		defaultUserTblModel: newUserTblModel(conn),
	}
}

func (m *customUserTblModel) withSession(session sqlx.Session) UserTblModel {
	return NewUserTblModel(sqlx.NewSqlConnFromSession(session))
}
