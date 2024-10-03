// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	billDetailTblFieldNames          = builder.RawFieldNames(&BillDetailTbl{})
	billDetailTblRows                = strings.Join(billDetailTblFieldNames, ",")
	billDetailTblRowsExpectAutoSet   = strings.Join(stringx.Remove(billDetailTblFieldNames), ",")
	billDetailTblRowsWithPlaceHolder = strings.Join(stringx.Remove(billDetailTblFieldNames, "`id`"), "=?,") + "=?"
)

type (
	billDetailTblModel interface {
		Insert(ctx context.Context, data *BillDetailTbl) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*BillDetailTbl, error)
		Update(ctx context.Context, data *BillDetailTbl) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBillDetailTblModel struct {
		conn  sqlx.SqlConn
		table string
	}

	BillDetailTbl struct {
		Id       int64          `db:"id"`
		BillId   sql.NullInt64  `db:"bill_id"`
		Name     sql.NullString `db:"name"`
		Price    sql.NullInt64  `db:"price"`
		Type     sql.NullInt64  `db:"type"`
		Quantity sql.NullInt64  `db:"quantity"`
	}
)

func newBillDetailTblModel(conn sqlx.SqlConn) *defaultBillDetailTblModel {
	return &defaultBillDetailTblModel{
		conn:  conn,
		table: "`bill_detail_tbl`",
	}
}

func (m *defaultBillDetailTblModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultBillDetailTblModel) FindOne(ctx context.Context, id int64) (*BillDetailTbl, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", billDetailTblRows, m.table)
	var resp BillDetailTbl
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBillDetailTblModel) Insert(ctx context.Context, data *BillDetailTbl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, billDetailTblRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.BillId, data.Name, data.Price, data.Type, data.Quantity)
	return ret, err
}

func (m *defaultBillDetailTblModel) Update(ctx context.Context, data *BillDetailTbl) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, billDetailTblRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.BillId, data.Name, data.Price, data.Type, data.Quantity, data.Id)
	return err
}

func (m *defaultBillDetailTblModel) tableName() string {
	return m.table
}
