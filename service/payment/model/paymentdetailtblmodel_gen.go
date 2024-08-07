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
		Id        int64  `db:"id"`
		PaymentId int64  `db:"payment_id"`
		Amount    int64  `db:"amount"`
		Type      int64  `db:"type"`
		Utl       string `db:"utl"`
		CreatedAt int64  `db:"created_at"`
		UpdatedAt int64  `db:"updated_at"`
		CreatedBy int64  `db:"created_by"`
		UpdatedBy int64  `db:"updated_by"`
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
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPaymentDetailTblModel) Insert(ctx context.Context, data *PaymentDetailTbl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, paymentDetailTblRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.PaymentId, data.Amount, data.Type, data.Utl, data.CreatedAt, data.UpdatedAt, data.CreatedBy, data.UpdatedBy)
	return ret, err
}

func (m *defaultPaymentDetailTblModel) Update(ctx context.Context, data *PaymentDetailTbl) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, paymentDetailTblRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.PaymentId, data.Amount, data.Type, data.Utl, data.CreatedAt, data.UpdatedAt, data.CreatedBy, data.UpdatedBy, data.Id)
	return err
}

func (m *defaultPaymentDetailTblModel) tableName() string {
	return m.table
}
