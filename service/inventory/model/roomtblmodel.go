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
		FindByHouseID(ctx context.Context, houseID, limit, offset int64) ([]*RoomTbl, int, error)
		DeleteByHouseID(ctx context.Context, houseID int64) error
		CountRoom(ctx context.Context, search string, houseType int64) (int, error)
		FilterRoom(ctx context.Context, search string, houseType, limit, offset int64) ([]*RoomTbl, error)
		FindByIDs(ctx context.Context, roomIDs []int64) ([]*RoomTbl, error)
		FindMultiByHouseIDs(ctx context.Context, houseIDs []int64) ([]*RoomTbl, error)
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

func (m *customRoomTblModel) FindByHouseID(ctx context.Context, houseID, limit, offset int64) ([]*RoomTbl, int, error) {
	query := fmt.Sprintf("select %s from %s where `house_id` = ? limit ? offset ?", roomTblRows, m.table)
	count := fmt.Sprintf("select count(*) from %s where `house_id` = ?", m.table)
	var resp []*RoomTbl
	var total int
	err := m.conn.QueryRowCtx(ctx, &total, count, houseID)
	if err != nil {
		return nil, 0, err
	}
	err = m.conn.QueryRowsCtx(ctx, &resp, query, houseID, limit, offset)
	switch err {
	case nil:
		return resp, total, nil
	case sqlx.ErrNotFound:
		return nil, 0, nil
	default:
		return nil, 0, err
	}
}

func (m *customRoomTblModel) DeleteByHouseID(ctx context.Context, houseID int64) error {
	query := fmt.Sprintf("delete from %s where `house_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, houseID)
	return err
}

func (m *customRoomTblModel) FilterRoom(ctx context.Context, search string, houseType, limit, offset int64) ([]*RoomTbl, error) {
	query := fmt.Sprintf("select %s from %s where `name` like ? ", roomTblRows, m.table)
	var vals []interface{}
	vals = append(vals, "%"+search+"%")

	if houseType != 0 {
		query += " and `house_id` in (select `id` from `house_tbl` where `type` = ?)"
		vals = append(vals, houseType)
	}
	if limit > 0 {
		query += " limit ? offset ?"
		vals = append(vals, limit, offset) 
	}

	var resp []*RoomTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, vals...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customRoomTblModel) CountRoom(ctx context.Context, search string, houseType int64) (int, error) {
	query := fmt.Sprintf("select count(*) from %s where `name` like ? ", m.table)
	var vals []interface{}
	vals = append(vals, "%"+search+"%")

	if houseType != 0 {
		query += " and `house_id` in (select `id` from `house_tbl` where `type` = ?)"
		vals = append(vals, houseType)
	}

	var total int
	err := m.conn.QueryRowCtx(ctx, &total, query, vals...)
	return total, err
}

func (m *customRoomTblModel) FindByIDs(ctx context.Context, roomIDs []int64) ([]*RoomTbl, error) {
	query := fmt.Sprintf("select %s from %s where `id` in (", roomTblRows, m.table)
    var resp []*RoomTbl
	var vals []interface{}
	for _, id := range roomIDs {
        query += "?,"
        vals = append(vals, id)
    }
	query = query[:len(query)-1] + ")"
    err := m.conn.QueryRowsCtx(ctx, &resp, query, vals...)
    switch err {
    case nil:
        return resp, nil
    case sqlx.ErrNotFound:
        return nil, nil
    default:
        return nil, err
    }
}

func (m *customRoomTblModel) FindMultiByHouseIDs(ctx context.Context, houseIDs []int64) ([]*RoomTbl, error) {
	query := fmt.Sprintf("select %s from %s where `house_id` in (", roomTblRows, m.table)
	var resp []*RoomTbl
	var vals []interface{}
	for _, id := range houseIDs {
		query += "?,"
		vals = append(vals, id)
	}
	query = query[:len(query)-1] + ")"
	err := m.conn.QueryRowsCtx(ctx, &resp, query, vals...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}