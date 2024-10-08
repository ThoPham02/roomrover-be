package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HouseTblModel = (*customHouseTblModel)(nil)

type (
	// HouseTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHouseTblModel.
	HouseTblModel interface {
		houseTblModel
		withSession(session sqlx.Session) HouseTblModel
		FilterHouse(ctx context.Context, userID int64, search string, limit, offset int64) (total int64, listHouses []*HouseTbl, err error)
		FindMultiByID(ctx context.Context, ids []int64) ([]*HouseTbl, error)
	}

	customHouseTblModel struct {
		*defaultHouseTblModel
	}
)

// NewHouseTblModel returns a model for the database table.
func NewHouseTblModel(conn sqlx.SqlConn) HouseTblModel {
	return &customHouseTblModel{
		defaultHouseTblModel: newHouseTblModel(conn),
	}
}

func (m *customHouseTblModel) withSession(session sqlx.Session) HouseTblModel {
	return NewHouseTblModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customHouseTblModel) FilterHouse(ctx context.Context, userID int64, search string, limit, offset int64) (total int64, listHouses []*HouseTbl, err error) {
	var searchVal string = "%" + search + "%"
	var vals []interface{}
	selectQuery := fmt.Sprintf("select %s from %s where `user_id` = ? and `name` like ?", houseTblRows, m.table)
	vals = append(vals, userID, searchVal)
	if limit > 0 {
		selectQuery += " limit ? offset ?"
		vals = append(vals, limit, offset)
	}
	err = m.conn.QueryRowsCtx(ctx, &listHouses, selectQuery, vals...)
	if err != nil {
		return 0, nil, err
	}
	countQuery := fmt.Sprintf("select count(*) from %s where `user_id` = ? and `name` like ?", m.table)
	err = m.conn.QueryRowCtx(ctx, &total, countQuery, userID, searchVal)
	if err != nil {
		return 0, nil, err
	}

	return total, listHouses, nil
}
func (m *customHouseTblModel) FindMultiByID(ctx context.Context, ids []int64) ([]*HouseTbl, error) {
	var listHouse []*HouseTbl
	var vals []interface{}
	query := fmt.Sprintf("select %s from %s where `id` in (", houseTblRows, m.table)
	for id := range ids {
		query += "?,"
		vals = append(vals, id)
	}
	query = query[:len(query)-1] + ")"

	err := m.conn.QueryRowsCtx(ctx, &listHouse, query, vals...)
	if err != nil {
		return nil, err
	}
	return listHouse, nil
}
