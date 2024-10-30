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
	paymentDetailTblFieldNames          = builder.RawFieldNames(&PaymentDetailTbl{})
	paymentDetailTblRows                = strings.Join(paymentDetailTblFieldNames, ",")
	paymentDetailTblRowsExpectAutoSet   = strings.Join(stringx.Remove(paymentDetailTblFieldNames), ",")
	paymentDetailTblRowsWithPlaceHolder = strings.Join(stringx.Remove(paymentDetailTblFieldNames, "`id`"), "=?,") + "=?"
)

type (
	paymentDetailTblModel interface {
		Insert(ctx context.Context, data *PaymentDetailTbl) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PaymentDetailTbl, error)
		Update(ctx context.Context, data *PaymentDetailTbl) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPaymentDetailTblModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PaymentDetailTbl struct {
		Id        int64          `db:"id"`
		PaymentId sql.NullInt64  `db:"payment_id"`
		Name      sql.NullString `db:"name"`
		Type      sql.NullInt64  `db:"type"`
		Price     sql.NullInt64  `db:"price"`
		Index     sql.NullInt64  `db:"index"`
	}
)

func newPaymentDetailTblModel(conn sqlx.SqlConn) *defaultPaymentDetailTblModel {
	return &defaultPaymentDetailTblModel{
		conn:  conn,
		table: "`payment_detail_tbl`",
	}
}

func (m *defaultPaymentDetailTblModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPaymentDetailTblModel) FindOne(ctx context.Context, id int64) (*PaymentDetailTbl, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", paymentDetailTblRows, m.table)
	var resp PaymentDetailTbl
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

func (m *defaultPaymentDetailTblModel) Insert(ctx context.Context, data *PaymentDetailTbl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, paymentDetailTblRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.PaymentId, data.Name, data.Type, data.Price, data.Index)
	return ret, err
}

func (m *defaultPaymentDetailTblModel) Update(ctx context.Context, data *PaymentDetailTbl) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, paymentDetailTblRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.PaymentId, data.Name, data.Type, data.Price, data.Index, data.Id)
	return err
}

func (m *defaultPaymentDetailTblModel) tableName() string {
	return m.table
}
