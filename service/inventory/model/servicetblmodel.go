package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ServiceTblModel = (*customServiceTblModel)(nil)

type (
	// ServiceTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customServiceTblModel.
	ServiceTblModel interface {
		serviceTblModel
		withSession(session sqlx.Session) ServiceTblModel
	}

	customServiceTblModel struct {
		*defaultServiceTblModel
	}
)

// NewServiceTblModel returns a model for the database table.
func NewServiceTblModel(conn sqlx.SqlConn) ServiceTblModel {
	return &customServiceTblModel{
		defaultServiceTblModel: newServiceTblModel(conn),
	}
}

func (m *customServiceTblModel) withSession(session sqlx.Session) ServiceTblModel {
	return NewServiceTblModel(sqlx.NewSqlConnFromSession(session))
}
