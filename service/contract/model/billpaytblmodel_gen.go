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
	billPayTblFieldNames          = builder.RawFieldNames(&BillPayTbl{})
	billPayTblRows                = strings.Join(billPayTblFieldNames, ",")
	billPayTblRowsExpectAutoSet   = strings.Join(stringx.Remove(billPayTblFieldNames), ",")
	billPayTblRowsWithPlaceHolder = strings.Join(stringx.Remove(billPayTblFieldNames, "`id`"), "=?,") + "=?"
)

type (
	billPayTblModel interface {
		Insert(ctx context.Context, data *BillPayTbl) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*BillPayTbl, error)
		Update(ctx context.Context, data *BillPayTbl) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBillPayTblModel struct {
		conn  sqlx.SqlConn
		table string
	}

	BillPayTbl struct {
		Id      int64          `db:"id"`
		BillId  int64          `db:"bill_id"`
		UserId  int64          `db:"user_id"`
		Amount  int64          `db:"amount"`
		PayDate int64          `db:"pay_date"`
		Status  int64          `db:"status"`
		TransId sql.NullString `db:"trans_id"`
		Type    int64          `db:"type"`
		Url     sql.NullString `db:"url"`
	}
)

func newBillPayTblModel(conn sqlx.SqlConn) *defaultBillPayTblModel {
	return &defaultBillPayTblModel{
		conn:  conn,
		table: "`bill_pay_tbl`",
	}
}

func (m *defaultBillPayTblModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultBillPayTblModel) FindOne(ctx context.Context, id int64) (*BillPayTbl, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", billPayTblRows, m.table)
	var resp BillPayTbl
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

func (m *defaultBillPayTblModel) Insert(ctx context.Context, data *BillPayTbl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, billPayTblRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.BillId, data.UserId, data.Amount, data.PayDate, data.Status, data.TransId, data.Type, data.Url)
	return ret, err
}

func (m *defaultBillPayTblModel) Update(ctx context.Context, data *BillPayTbl) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, billPayTblRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.BillId, data.UserId, data.Amount, data.PayDate, data.Status, data.TransId, data.Type, data.Url, data.Id)
	return err
}

func (m *defaultBillPayTblModel) tableName() string {
	return m.table
}
