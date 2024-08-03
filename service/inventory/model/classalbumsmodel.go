package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ClassAlbumsModel = (*customClassAlbumsModel)(nil)

type (
	// ClassAlbumsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customClassAlbumsModel.
	ClassAlbumsModel interface {
		classAlbumsModel
		withSession(session sqlx.Session) ClassAlbumsModel
	}

	customClassAlbumsModel struct {
		*defaultClassAlbumsModel
	}
)

// NewClassAlbumsModel returns a model for the database table.
func NewClassAlbumsModel(conn sqlx.SqlConn) ClassAlbumsModel {
	return &customClassAlbumsModel{
		defaultClassAlbumsModel: newClassAlbumsModel(conn),
	}
}

func (m *customClassAlbumsModel) withSession(session sqlx.Session) ClassAlbumsModel {
	return NewClassAlbumsModel(sqlx.NewSqlConnFromSession(session))
}
