// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	billDetailFieldNames          = builder.RawFieldNames(&BillDetail{})
	billDetailRows                = strings.Join(billDetailFieldNames, ",")
	billDetailRowsExpectAutoSet   = strings.Join(stringx.Remove(billDetailFieldNames), ",")
	billDetailRowsWithPlaceHolder = strings.Join(stringx.Remove(billDetailFieldNames, "`id`"), "=?,") + "=?"
)

type (
	billDetailModel interface {
		Insert(ctx context.Context, data *BillDetail) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*BillDetail, error)
		Update(ctx context.Context, data *BillDetail) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBillDetailModel struct {
		conn  sqlx.SqlConn
		table string
	}

	BillDetail struct {
		Id                int64 `db:"id"`
		BillId            int64 `db:"bill_id"`
		ContractServiceId int64 `db:"contract_service_id"`
		Price             int64 `db:"price"`
		Type              int64 `db:"type"`
		Quantity          int64 `db:"quantity"`
	}
)

func newBillDetailModel(conn sqlx.SqlConn) *defaultBillDetailModel {
	return &defaultBillDetailModel{
		conn:  conn,
		table: "`bill_detail`",
	}
}

func (m *defaultBillDetailModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultBillDetailModel) FindOne(ctx context.Context, id int64) (*BillDetail, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", billDetailRows, m.table)
	var resp BillDetail
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBillDetailModel) Insert(ctx context.Context, data *BillDetail) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, billDetailRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.BillId, data.ContractServiceId, data.Price, data.Type, data.Quantity)
	return ret, err
}

func (m *defaultBillDetailModel) Update(ctx context.Context, data *BillDetail) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, billDetailRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.BillId, data.ContractServiceId, data.Price, data.Type, data.Quantity, data.Id)
	return err
}

func (m *defaultBillDetailModel) tableName() string {
	return m.table
}
