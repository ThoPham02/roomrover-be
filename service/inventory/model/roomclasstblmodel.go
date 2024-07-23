package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RoomClassTblModel = (*customRoomClassTblModel)(nil)

type (
	// RoomClassTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoomClassTblModel.
	RoomClassTblModel interface {
		roomClassTblModel
		withSession(session sqlx.Session) RoomClassTblModel
	}

	customRoomClassTblModel struct {
		*defaultRoomClassTblModel
	}
)

// NewRoomClassTblModel returns a model for the database table.
func NewRoomClassTblModel(conn sqlx.SqlConn) RoomClassTblModel {
	return &customRoomClassTblModel{
		defaultRoomClassTblModel: newRoomClassTblModel(conn),
	}
}

func (m *customRoomClassTblModel) withSession(session sqlx.Session) RoomClassTblModel {
	return NewRoomClassTblModel(sqlx.NewSqlConnFromSession(session))
}
