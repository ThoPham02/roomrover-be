package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RoomTblModel = (*customRoomTblModel)(nil)

type (
	// RoomTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoomTblModel.
	RoomTblModel interface {
		roomTblModel
		withSession(session sqlx.Session) RoomTblModel
	}

	customRoomTblModel struct {
		*defaultRoomTblModel
	}
)

// NewRoomTblModel returns a model for the database table.
func NewRoomTblModel(conn sqlx.SqlConn) RoomTblModel {
	return &customRoomTblModel{
		defaultRoomTblModel: newRoomTblModel(conn),
	}
}

func (m *customRoomTblModel) withSession(session sqlx.Session) RoomTblModel {
	return NewRoomTblModel(sqlx.NewSqlConnFromSession(session))
}
