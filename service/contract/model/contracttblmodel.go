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
		withSession(session sqlx.Session) ContractTblModel
		GetContractByRoomID(ctx context.Context, roomID int64) (*ContractTbl, error)
		GetContractByTime(ctx context.Context, time int64) ([]*ContractTbl, error)
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

func (m *customContractTblModel) withSession(session sqlx.Session) ContractTblModel {
	return NewContractTblModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customContractTblModel) GetContractByRoomID(ctx context.Context, roomID int64) (*ContractTbl, error) {
	query := fmt.Sprintf("select %s from %s where `room_id` = ? limit 1", contractTblRows, m.table)
	var resp ContractTbl
	err := m.conn.QueryRowCtx(ctx, &resp, query, roomID)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customContractTblModel) GetContractByTime(ctx context.Context, time int64) ([]*ContractTbl, error) {
	var startTime = time - 2*86400000 // lay ra 2 ngay truoc
	query := fmt.Sprintf("select %s from %s where `next_bill` between ? and ?", contractTblRows, m.table)
	var resp []*ContractTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, startTime, time)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customContractTblModel) CountContractByCondition(ctx context.Context, renterID int64, lessorID int64, search string, status int64, createFrom int64, createTo int64) (int64, error) {
	var query string
	var resp int64
	var err error
	var vals []interface{}
	query = fmt.Sprintf("select count(*) from %s where `code` like ? ", m.table)
	vals = append(vals, "%"+search+"%")
	if status != 0 {
		query += " and `status` = ?"
		vals = append(vals, status)
	}
	query += " and `created_at` between ? and ? "
	vals = append(vals, createFrom, createTo)
	if renterID != 0 {
		query += " and `renter_id` = ?"
		vals = append(vals, renterID)
	}
	if lessorID != 0 {
		query += " and `lessor_id` = ?"
		vals = append(vals, lessorID)
	}

	err = m.conn.QueryRowCtx(ctx, &resp, query, vals...)
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
	var resp []*ContractTbl
	var err error
	var vals []interface{}
	query = fmt.Sprintf("select %s from %s where `code` like ? ",contractTblRows, m.table)
	vals = append(vals, "%"+search+"%")
	if status != 0 {
		query += " and `status` = ?"
		vals = append(vals, status)
	}
	query += " and `created_at` between ? and ? "
	vals = append(vals, createFrom, createTo)
	if renterID != 0 {
		query += " and `renter_id` = ?"
		vals = append(vals, renterID)
	}
	if lessorID != 0 {
		query += " and `lessor_id` = ?"
		vals = append(vals, lessorID)
	}
	if limit > 0 {
		query += " limit ? offset ?"
		vals = append(vals, limit, offset)
	}
	err = m.conn.QueryRowsCtx(ctx, &resp, query, vals...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
