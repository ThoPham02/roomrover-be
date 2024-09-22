package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RoomTblModel = (*customRoomTblModel)(nil)

type (
	// RoomTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoomTblModel.
	RoomTblModel interface {
		roomTblModel
		withSession(session sqlx.Session) RoomTblModel
		FindByHouseID(ctx context.Context, houseID int64) ([]*RoomTbl, error)
		DeleteByHouseID(ctx context.Context, houseID int64) error
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

func (m *customRoomTblModel) FindByHouseID(ctx context.Context, houseID int64) ([]*RoomTbl, error) {
	query := fmt.Sprintf("select %s from %s where `house_id` = ?", roomTblRows, m.table)
	var resp []*RoomTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, houseID)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customRoomTblModel) DeleteByHouseID(ctx context.Context, houseID int64) error {
	query := fmt.Sprintf("delete from %s where `house_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, houseID)
	return err
}
