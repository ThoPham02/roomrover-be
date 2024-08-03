package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ DistrictsModel = (*customDistrictsModel)(nil)

type (
	// DistrictsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDistrictsModel.
	DistrictsModel interface {
		districtsModel
		withSession(session sqlx.Session) DistrictsModel
	}

	customDistrictsModel struct {
		*defaultDistrictsModel
	}
)

// NewDistrictsModel returns a model for the database table.
func NewDistrictsModel(conn sqlx.SqlConn) DistrictsModel {
	return &customDistrictsModel{
		defaultDistrictsModel: newDistrictsModel(conn),
	}
}

func (m *customDistrictsModel) withSession(session sqlx.Session) DistrictsModel {
	return NewDistrictsModel(sqlx.NewSqlConnFromSession(session))
}
