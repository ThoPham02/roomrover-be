package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RoomAlbumsTblModel = (*customRoomAlbumsTblModel)(nil)

type (
	// RoomAlbumsTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoomAlbumsTblModel.
	RoomAlbumsTblModel interface {
		roomAlbumsTblModel
		withSession(session sqlx.Session) RoomAlbumsTblModel
	}

	customRoomAlbumsTblModel struct {
		*defaultRoomAlbumsTblModel
	}
)

// NewRoomAlbumsTblModel returns a model for the database table.
func NewRoomAlbumsTblModel(conn sqlx.SqlConn) RoomAlbumsTblModel {
	return &customRoomAlbumsTblModel{
		defaultRoomAlbumsTblModel: newRoomAlbumsTblModel(conn),
	}
}

func (m *customRoomAlbumsTblModel) withSession(session sqlx.Session) RoomAlbumsTblModel {
	return NewRoomAlbumsTblModel(sqlx.NewSqlConnFromSession(session))
}
