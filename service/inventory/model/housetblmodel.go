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
	selectQuery := fmt.Sprintf("select %s from %s where `name` like ?", houseTblRows, m.table)
	vals = append(vals, searchVal)
	if limit > 0 {
		selectQuery += " limit ? offset ?"
		vals = append(vals, limit, offset)
	}
	err = m.conn.QueryRowCtx(ctx, &total, selectQuery, vals...)
	if err != nil {
		return 0, nil, err
	}
	countQuery := fmt.Sprintf("select count(*) from %s where `name` like ?", m.table)
	err = m.conn.QueryRowsCtx(ctx, &listHouses, countQuery, searchVal)
	if err != nil {
		return 0, nil, err
	}

	return total, listHouses, nil
}
