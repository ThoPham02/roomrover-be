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
	contactTblFieldNames          = builder.RawFieldNames(&ContactTbl{})
	contactTblRows                = strings.Join(contactTblFieldNames, ",")
	contactTblRowsExpectAutoSet   = strings.Join(stringx.Remove(contactTblFieldNames), ",")
	contactTblRowsWithPlaceHolder = strings.Join(stringx.Remove(contactTblFieldNames, "`id`"), "=?,") + "=?"
)

type (
	contactTblModel interface {
		Insert(ctx context.Context, data *ContactTbl) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ContactTbl, error)
		Update(ctx context.Context, data *ContactTbl) error
		Delete(ctx context.Context, id int64) error
	}

	defaultContactTblModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ContactTbl struct {
		Id       int64         `db:"id"`
		HouseId  sql.NullInt64 `db:"house_id"`
		RenterId sql.NullInt64 `db:"renter_id"`
		LessorId sql.NullInt64 `db:"lessor_id"`
		Datetime sql.NullInt64 `db:"datetime"`
		Status   int64         `db:"status"`
	}
)

func newContactTblModel(conn sqlx.SqlConn) *defaultContactTblModel {
	return &defaultContactTblModel{
		conn:  conn,
		table: "`contact_tbl`",
	}
}

func (m *defaultContactTblModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultContactTblModel) FindOne(ctx context.Context, id int64) (*ContactTbl, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", contactTblRows, m.table)
	var resp ContactTbl
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

func (m *defaultContactTblModel) Insert(ctx context.Context, data *ContactTbl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, contactTblRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.HouseId, data.RenterId, data.LessorId, data.Datetime, data.Status)
	return ret, err
}

func (m *defaultContactTblModel) Update(ctx context.Context, data *ContactTbl) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, contactTblRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.HouseId, data.RenterId, data.LessorId, data.Datetime, data.Status, data.Id)
	return err
}

func (m *defaultContactTblModel) tableName() string {
	return m.table
}
