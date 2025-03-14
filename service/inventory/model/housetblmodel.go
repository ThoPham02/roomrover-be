package model

import (
	"context"
	"fmt"
	"roomrover/common"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HouseTblModel = (*customHouseTblModel)(nil)

type (
	// HouseTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHouseTblModel.
	HouseTblModel interface {
		houseTblModel
		withSession(session sqlx.Session) HouseTblModel
		FilterHouse(ctx context.Context, userID int64, search string, houseType, status, limit, offset int64) (total int64, listHouses []*HouseTbl, err error)
		FindMultiByID(ctx context.Context, ids []int64) ([]*HouseTbl, error)
		SearchHouse(ctx context.Context, search string, districtID, provinceID, wardID, typeHouse, unit int64, priceFrom, priceTo, areaFrom, areaTo, limit, offset int64) (total int64, listHouses []*HouseTbl, err error)
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

func (m *customHouseTblModel) FilterHouse(ctx context.Context, userID int64, search string, houseType, status, limit, offset int64) (total int64, listHouses []*HouseTbl, err error) {
	var searchVal string = "%" + search + "%"
	var vals []interface{}
	selectQuery := fmt.Sprintf("select %s from %s where `user_id` = ? and `name` like ?", houseTblRows, m.table)
	vals = append(vals, userID, searchVal)
	if houseType > 0 {
		selectQuery += " and `type` = ?"
		vals = append(vals, houseType)
	}
	if status > 0 {
		selectQuery += " and `status` = ?"
		vals = append(vals, status)
	}
	selectQuery += " order by `updated_at` desc"

	countQuery := fmt.Sprintf("select count(*) from %s where `user_id` = ? and `name` like ?", m.table)
	if houseType > 0 {
		countQuery += " and `type` = ?"
	}
	if status > 0 {
		countQuery += " and `status` = ?"
	}
	err = m.conn.QueryRowCtx(ctx, &total, countQuery, vals...)
	if err != nil {
		return 0, nil, err
	}

	if limit > 0 {
		selectQuery += " limit ? offset ?"
		vals = append(vals, limit, offset)
	}
	err = m.conn.QueryRowsCtx(ctx, &listHouses, selectQuery, vals...)
	if err != nil {
		return 0, nil, err
	}

	return total, listHouses, nil
}
func (m *customHouseTblModel) FindMultiByID(ctx context.Context, ids []int64) ([]*HouseTbl, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var listHouse []*HouseTbl
	var vals []interface{}
	query := fmt.Sprintf("select %s from %s where `id` in (", houseTblRows, m.table)
	for _, id := range ids {
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

func (m *customHouseTblModel) SearchHouse(ctx context.Context, search string, districtID, provinceID, wardID, typeHouse, unit, priceFrom, priceTo, areaFrom, areaTo, limit, offset int64) (total int64, listHouses []*HouseTbl, err error) {
	var searchVal string = "%" + search + "%"
	var vals []interface{}
	selectQuery := fmt.Sprintf("select %s from %s where `status` = ? and `name` like ?", houseTblRows, m.table)
	countQuery := fmt.Sprintf("select count(*) from %s where `status` = ? and `name` like ?", m.table)
	vals = append(vals, common.HOUSE_STATUS_ACTIVE, searchVal)
	if districtID > 0 {
		countQuery += " and `district_id` = ?"
		selectQuery += " and `district_id` = ?"
		vals = append(vals, districtID)
	}
	if provinceID > 0 {
		countQuery += " and `province_id` = ?"
		selectQuery += " and `province_id` = ?"
		vals = append(vals, provinceID)
	}
	if wardID > 0 {
		countQuery += " and `ward_id` = ?"
		selectQuery += " and `ward_id` = ?"
		vals = append(vals, wardID)
	}
	if priceFrom > 0 {
		countQuery += " and `price` >= ?"
		selectQuery += " and `price` >= ?"
		vals = append(vals, priceFrom)
	}
	if priceTo > 0 {
		countQuery += " and `price` <= ?"
		selectQuery += " and `price` <= ?"
		vals = append(vals, priceTo)
	}
	if areaFrom > 0 {
		countQuery += " and `area` >= ?"
		selectQuery += " and `area` >= ?"
		vals = append(vals, areaFrom)
	}
	if areaTo > 0 {
		countQuery += " and `area` <= ?"
		selectQuery += " and `area` <= ?"
		vals = append(vals, areaTo)
	}
	if typeHouse > 0 {
		countQuery += " and `type` = ?"
		selectQuery += " and `type` = ?"
		vals = append(vals, typeHouse)
	}
	if unit > 0 {
		countQuery += " and `unit` & ?"
		selectQuery += " and `unit` & ?"
		vals = append(vals, unit)
	}

	err = m.conn.QueryRowCtx(ctx, &total, countQuery, vals...)
	if err != nil {
		return 0, nil, err
	}
	selectQuery += " order by `updated_at` desc"
	if limit > 0 {
		selectQuery += " limit ? offset ?"
		vals = append(vals, limit, offset)
	}
	err = m.conn.QueryRowsCtx(ctx, &listHouses, selectQuery, vals...)
	if err != nil {
		return 0, nil, err
	}

	return total, listHouses, nil
}
