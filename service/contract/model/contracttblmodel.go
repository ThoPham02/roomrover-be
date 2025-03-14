package model

import (
	"context"
	"fmt"
	"roomrover/common"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContractTblModel = (*customContractTblModel)(nil)

type (
	// ContractTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContractTblModel.
	ContractTblModel interface {
		contractTblModel
		CountByHouseID(ctx context.Context, houseID int64) (int64, error)
		CountContractByCondition(ctx context.Context, lessorID int64, renterID int64, search string, status int64, createFrom int64, createTo int64) (int64, error)
		FindContractByCondition(ctx context.Context, lessorID int64, renterID int64, search string, status int64, createFrom int64, createTo int64, offset int64, limit int64) ([]*ContractTbl, error)
		FindActiveByRoomID(ctx context.Context, roomID int64) (*ContractTbl, error)
		FilterContractOutDate(ctx context.Context, checkTime int64) ([]*ContractTbl, error)
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

func (m *customContractTblModel) CountContractByCondition(ctx context.Context, lessorID int64, renterID int64, search string, status int64, createFrom int64, createTo int64) (int64, error) {
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
		query += " and `created_at` >= ?"
		args = append(args, createFrom)
	}
	if createTo != 0 {
		query += " and `created_at` <= ?"
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

func (m *customContractTblModel) FindContractByCondition(ctx context.Context, lessorID int64, renterID int64, search string, status int64, createFrom int64, createTo int64, offset int64, limit int64) ([]*ContractTbl, error) {
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
		query += " and `created_at` >= ?"
		args = append(args, createFrom)
	}
	if createTo != 0 {
		query += " and `created_at` <= ?"
		args = append(args, createTo)
	}
	query = fmt.Sprintf("select %s from %s where 1=1 %s", contractTblRows, m.table, query)
	query += " order by `updated_at` desc"
	if limit != 0 {
		query += " limit ? offset ?"
		args = append(args, limit, offset)
	}
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

func (m *customContractTblModel) FindActiveByRoomID(ctx context.Context, roomID int64) (*ContractTbl, error) {
	query := fmt.Sprintf("select %s from %s where `room_id` = ? and `status` & 1", contractTblRows, m.table)
	var resp ContractTbl
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

func (m *customContractTblModel) FilterContractOutDate(ctx context.Context, checkTime int64) ([]*ContractTbl, error) {
	query := fmt.Sprintf("select %s from %s where `id` in (select `contract_id` from `payment_tbl` where `deposit_date` between ? and ?) and `status` = ?", contractTblRows, m.table)
	var resp []*ContractTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, checkTime-43200000, checkTime+43200000, common.CONTRACT_STATUS_WAIT_DEPOSIT)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
