package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RoomTblModel = (*customRoomTblModel)(nil)

type (
	// RoomTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoomTblModel.
	HouseRoomTbl struct {
		Id            int64          `db:"id"`
		HouseId       sql.NullInt64  `db:"house_id"`
		HouseRoomName sql.NullString `db:"house_room_name"`
		DistrictID    sql.NullInt64  `db:"district_id"`
		ProvinceID    sql.NullInt64  `db:"province_id"`
		WardID        sql.NullInt64  `db:"ward_id"`
		Address       sql.NullString `db:"address"`
		Area          sql.NullInt64  `db:"area"`
		Type          sql.NullInt64  `db:"type"`
		Price         sql.NullInt64  `db:"price"`
		Status        int64          `db:"status"`
		Capacity      sql.NullInt64  `db:"capacity"`
		EIndex        sql.NullInt64  `db:"e_index"`
		WIndex        sql.NullInt64  `db:"w_index"`
	}

	RoomTblModel interface {
		roomTblModel
		withSession(session sqlx.Session) RoomTblModel
		FindByHouseID(ctx context.Context, houseID, limit, offset int64) ([]*RoomTbl, int, error)
		DeleteByHouseID(ctx context.Context, houseID int64) error
		CountRoom(ctx context.Context, userID int64, search string, houseType, status int64) (int, error)
		FilterRoom(ctx context.Context, userID int64, search string, houseType, status, limit, offset int64) ([]*RoomTbl, error)
		FindByIDs(ctx context.Context, roomIDs []int64) ([]*RoomTbl, error)
		FindMultiByHouseIDs(ctx context.Context, houseIDs []int64) ([]*RoomTbl, error)
		SearchRoom(ctx context.Context, userID, houseType int64, search string, status int64, limit, offset int64) ([]*HouseRoomTbl, error)
		CountSearchRoom(ctx context.Context, userID, houseType int64, search string, status int64) (int, error)
		GetHouseRoomByRoomID(ctx context.Context, roomID int64) (*HouseRoomTbl, error)
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
	var vals []interface{}
	query := fmt.Sprintf("select %s from %s where `house_id` = ? ", roomTblRows, m.table)
	vals = append(vals, houseID)
	if limit > 0 {
		query += "limit ? offset ?"
		vals = append(vals, limit, offset)
	}
	count := fmt.Sprintf("select count(*) from %s where `house_id` = ?", m.table)
	var resp []*RoomTbl
	var total int
	err := m.conn.QueryRowCtx(ctx, &total, count, houseID)
	if err != nil {
		return nil, 0, err
	}
	err = m.conn.QueryRowsCtx(ctx, &resp, query, vals...)
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

func (m *customRoomTblModel) FilterRoom(ctx context.Context, userID int64, search string, houseType, status, limit, offset int64) ([]*RoomTbl, error) {
	query := fmt.Sprintf("select %s from %s where `name` like ? ", roomTblRows, m.table)
	var vals []interface{}
	vals = append(vals, "%"+search+"%")

	if houseType != 0 {
		query += " and `house_id` in (select `id` from `house_tbl` where `type` = ? and `user_id` = ?)"
		vals = append(vals, houseType, userID)
	} else {
		query += " and `house_id` in (select `id` from `house_tbl` where `user_id` = ?)"
		vals = append(vals, userID)
	}

	if status != 0 {
		query += " and `status` = ?"
		vals = append(vals, status)
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

func (m *customRoomTblModel) CountRoom(ctx context.Context, userID int64, search string, houseType, status int64) (int, error) {
	query := fmt.Sprintf("select count(*) from %s where `name` like ? ", m.table)
	var vals []interface{}
	vals = append(vals, "%"+search+"%")

	if houseType != 0 {
		query += " and `house_id` in (select `id` from `house_tbl` where `type` = ? and `user_id` = ?)"
		vals = append(vals, houseType, userID)
	} else {
		query += " and `house_id` in (select `id` from `house_tbl` where `user_id` = ?)"
		vals = append(vals, userID)
	}
	if status != 0 {
		query += " and `status` = ?"
		vals = append(vals, status)
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

func (m *customRoomTblModel) SearchRoom(ctx context.Context, userID, houseType int64, search string, status int64, limit, offset int64) ([]*HouseRoomTbl, error) {
	query := `SELECT
    room_tbl.id AS id,
    house_tbl.id AS house_id,
    CONCAT(house_tbl.name, ' - ', room_tbl.name) AS house_room_name,
    house_tbl.district_id AS district_id,
    house_tbl.province_id AS province_id,
    house_tbl.ward_id AS ward_id,
    house_tbl.address AS address,
    house_tbl.area AS area,
    house_tbl.price AS price,
	house_tbl.type AS type,
    room_tbl.status AS status,
    room_tbl.capacity AS capacity,
    room_tbl.e_index AS e_index,
    room_tbl.w_index AS w_index
	FROM room_tbl
	JOIN house_tbl ON room_tbl.house_id = house_tbl.id
	WHERE CONCAT(house_tbl.name, ' - ', room_tbl.name) LIKE ? 
	AND house_tbl.user_id = ?`
	var vals []interface{}
	vals = append(vals, "%"+search+"%", userID)

	if houseType != 0 {
		query += " and house_tbl.type = ? "
		vals = append(vals, houseType)
	}
	if status != 0 {
		query += " and room_tbl.status = ?"
		vals = append(vals, status)
	}

	if limit > 0 {
		query += " limit ? offset ?"
		vals = append(vals, limit, offset)
	}

	var resp []*HouseRoomTbl
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

func (m *customRoomTblModel) CountSearchRoom(ctx context.Context, userID, houseType int64, search string, status int64) (int, error) {
	query := `SELECT
	COUNT(*)
	FROM room_tbl
	JOIN house_tbl ON room_tbl.house_id = house_tbl.id
	WHERE CONCAT(house_tbl.name, ' - ', room_tbl.name) LIKE ? 
	AND house_tbl.user_id = ?`
	var vals []interface{}
	vals = append(vals, "%"+search+"%", userID)

	if houseType != 0 {
		query += " and house_tbl.type = ? "
		vals = append(vals, houseType)
	}
	if status != 0 {
		query += " and room_tbl.status = ?"
		vals = append(vals, status)
	}

	var total int
	err := m.conn.QueryRowCtx(ctx, &total, query, vals...)
	return total, err
}

func (m *customRoomTblModel) GetHouseRoomByRoomID(ctx context.Context, roomID int64) (*HouseRoomTbl, error) {
	query := `SELECT
    room_tbl.id AS id,
    house_tbl.id AS house_id,
    CONCAT(house_tbl.name, ' - ', room_tbl.name) AS house_room_name,
    house_tbl.district_id AS district_id,
    house_tbl.province_id AS province_id,
    house_tbl.ward_id AS ward_id,
    house_tbl.address AS address,
    house_tbl.area AS area,
    house_tbl.price AS price,
    house_tbl.type AS type,
    room_tbl.status AS status,
    room_tbl.capacity AS capacity,
    room_tbl.e_index AS e_index,
    room_tbl.w_index AS w_index
	FROM room_tbl
	JOIN house_tbl ON room_tbl.house_id = house_tbl.id
	WHERE room_tbl.id = ?`
	var resp HouseRoomTbl
	err := m.conn.QueryRowCtx(ctx, &resp, query, roomID)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
