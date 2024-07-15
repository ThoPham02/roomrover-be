package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RoomsTblModel = (*customRoomsTblModel)(nil)

type (
	// RoomsTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoomsTblModel.
	RoomsTblModel interface {
		roomsTblModel
		withSession(session sqlx.Session) RoomsTblModel
	}

	customRoomsTblModel struct {
		*defaultRoomsTblModel
	}
)

// NewRoomsTblModel returns a model for the database table.
func NewRoomsTblModel(conn sqlx.SqlConn) RoomsTblModel {
	return &customRoomsTblModel{
		defaultRoomsTblModel: newRoomsTblModel(conn),
	}
}

func (m *customRoomsTblModel) withSession(session sqlx.Session) RoomsTblModel {
	return NewRoomsTblModel(sqlx.NewSqlConnFromSession(session))
}
