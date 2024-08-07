package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AlbumTblModel = (*customAlbumTblModel)(nil)

type (
	// AlbumTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAlbumTblModel.
	AlbumTblModel interface {
		albumTblModel
		withSession(session sqlx.Session) AlbumTblModel
	}

	customAlbumTblModel struct {
		*defaultAlbumTblModel
	}
)

// NewAlbumTblModel returns a model for the database table.
func NewAlbumTblModel(conn sqlx.SqlConn) AlbumTblModel {
	return &customAlbumTblModel{
		defaultAlbumTblModel: newAlbumTblModel(conn),
	}
}

func (m *customAlbumTblModel) withSession(session sqlx.Session) AlbumTblModel {
	return NewAlbumTblModel(sqlx.NewSqlConnFromSession(session))
}
