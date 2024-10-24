package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BillTblModel = (*customBillTblModel)(nil)

type FilterCondition struct {
	Search     string
	RenterID   int64
	LessorID   int64
	CreateFrom int64
	CreateTo   int64
	Status     int64
	Limit      int64
	Offset     int64
}

type (
	// BillTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBillTblModel.
	BillTblModel interface {
		billTblModel
		CountByCondition(ctx context.Context, condition FilterCondition) (int64, error)
		FilterBillByCondition(ctx context.Context, condition FilterCondition) ([]*BillTbl, error)
	}

	customBillTblModel struct {
		*defaultBillTblModel
	}
)

// NewBillTblModel returns a model for the database table.
func NewBillTblModel(conn sqlx.SqlConn) BillTblModel {
	return &customBillTblModel{
		defaultBillTblModel: newBillTblModel(conn),
	}
}

func (m *customBillTblModel) CountByCondition(ctx context.Context, condition FilterCondition) (int64, error) {
	query := fmt.Sprintf("select count(*) from %s where 1=1 ", m.table)
	var args []interface{}
	var count int64
	if condition.Search != "" {
		query += " and (title like ?)"
		args = append(args, "%"+condition.Search+"%")
	}
	if condition.RenterID != 0 {
		query += " and `payment_id` in (select `id` from `payment_tbl` where `contract_id` in (select `id` from `contract_tbl` where `renter_id` = ?))"
		args = append(args, condition.RenterID)
	}
	if condition.LessorID != 0 {
		query += " and `payment_id` in (select `id` from `payment_tbl` where `contract_id` in (select `id` from `contract_tbl` where `lessor_id` = ?))"
		args = append(args, condition.LessorID)
	}
	if condition.Status != 0 {
		query += " and `status` = ?"
		args = append(args, condition.Status)
	}

	err := m.conn.QueryRowCtx(ctx, &count, query, args...)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *customBillTblModel) FilterBillByCondition(ctx context.Context, condition FilterCondition) ([]*BillTbl, error) {
	query := fmt.Sprintf("select %s from %s where 1=1 ", billTblRows, m.table)
	var args []interface{}
	var resp []*BillTbl
	if condition.Search != "" {
		query += " and (title like ?)"
		args = append(args, "%"+condition.Search+"%")
	}
	if condition.RenterID != 0 {
		query += " and `payment_id` in (select `id` from `payment_tbl` where `contract_id` in (select `id` from `contract_tbl` where `renter_id` = ?))"
		args = append(args, condition.RenterID)
	}
	if condition.LessorID != 0 {
		query += " and `payment_id` in (select `id` from `payment_tbl` where `contract_id` in (select `id` from `contract_tbl` where `lessor_id` = ?))"
		args = append(args, condition.LessorID)
	}
	if condition.Limit != 0 {
		query += " limit ? offset ?"
		args = append(args, condition.Limit, condition.Offset)
	}

	err := m.conn.QueryRowsCtx(ctx, &resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
