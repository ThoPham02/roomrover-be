package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContractTblModel = (*customContractTblModel)(nil)

type (
	// ContractTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContractTblModel.
	ContractTblModel interface {
		contractTblModel
		CountByHouseID(ctx context.Context, houseID int64) (int64, error)
		CountContractByCondition(ctx context.Context, renterID int64, lessorID int64, search string, status int64, createFrom int64, createTo int64) (int64, error)
		FindContractByCondition(ctx context.Context, renterID int64, lessorID int64, search string, status int64, createFrom int64, createTo int64, offset int64, limit int64) ([]*ContractTbl, error)
	}

	customContractTblModel struct {
		*defaultContractTblModel
	}
)

// NewContractTblModel returns a model for the database table.
func NewContractTblModel(conn sqlx.SqlConn) ContractTblModel {
	return &customContractTblModel{
		defaultContractTblModel: newContractTblModel(conn),
	}
}

func (m *customContractTblModel) CountByHouseID(ctx context.Context, houseID int64) (int64, error) {
	query := fmt.Sprintf("select count(*) from %s where `room_id` in (select `id` from `room_tbl` where `house_id` = ?)", m.table)
	var resp int64
	err := m.conn.QueryRowCtx(ctx, &resp, query, houseID)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}
}

func (m *customContractTblModel) CountContractByCondition(ctx context.Context, renterID int64, lessorID int64, search string, status int64, createFrom int64, createTo int64) (int64, error) {
	var query string
	var args []interface{}
	if renterID != 0 {
		query += " and `renter_id` = ?"
		args = append(args, renterID)
	}
	if lessorID != 0 {
		query += " and `lessor_id` = ?"
		args = append(args, lessorID)
	}
	if search != "" {
		query += " and `code` like ?"
		args = append(args, "%"+search+"%")
	}
	if status != 0 {
		query += " and `status` = ?"
		args = append(args, status)
	}
	if createFrom != 0 {
		query += " and `create_time` >= ?"
		args = append(args, createFrom)
	}
	if createTo != 0 {
		query += " and `create_time` <= ?"
		args = append(args, createTo)
	}
	query = fmt.Sprintf("select count(*) from %s where 1=1 %s", m.table, query)
	var resp int64
	err := m.conn.QueryRowCtx(ctx, &resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}
}

func (m *customContractTblModel) FindContractByCondition(ctx context.Context, renterID int64, lessorID int64, search string, status int64, createFrom int64, createTo int64, offset int64, limit int64) ([]*ContractTbl, error) {
	var query string
	var args []interface{}
	if renterID != 0 {
		query += " and `renter_id` = ?"
		args = append(args, renterID)
	}
	if lessorID != 0 {
		query += " and `lessor_id` = ?"
		args = append(args, lessorID)
	}
	if search != "" {
		query += " and `code` like ?"
		args = append(args, "%"+search+"%")
	}
	if status != 0 {
		query += " and `status` = ?"
		args = append(args, status)
	}
	if createFrom != 0 {
		query += " and `create_time` >= ?"
		args = append(args, createFrom)
	}
	if createTo != 0 {
		query += " and `create_time` <= ?"
		args = append(args, createTo)
	}
	if limit != 0 {
		query += " limit ? offset ?"
		args = append(args, limit, offset)
	}
	query = fmt.Sprintf("select %s from %s where 1=1 %s", contractTblRows, m.table, query)
	var resp []*ContractTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
