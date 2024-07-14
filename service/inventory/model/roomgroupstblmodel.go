package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RoomGroupsTblModel = (*customRoomGroupsTblModel)(nil)

type (
	// RoomGroupsTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoomGroupsTblModel.
	RoomGroupsTblModel interface {
		roomGroupsTblModel
		withSession(session sqlx.Session) RoomGroupsTblModel
	}

	customRoomGroupsTblModel struct {
		*defaultRoomGroupsTblModel
	}
)

// NewRoomGroupsTblModel returns a model for the database table.
func NewRoomGroupsTblModel(conn sqlx.SqlConn) RoomGroupsTblModel {
	return &customRoomGroupsTblModel{
		defaultRoomGroupsTblModel: newRoomGroupsTblModel(conn),
	}
}

func (m *customRoomGroupsTblModel) withSession(session sqlx.Session) RoomGroupsTblModel {
	return NewRoomGroupsTblModel(sqlx.NewSqlConnFromSession(session))
}
