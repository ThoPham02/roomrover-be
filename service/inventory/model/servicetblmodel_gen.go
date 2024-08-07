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
	serviceTblFieldNames          = builder.RawFieldNames(&ServiceTbl{})
	serviceTblRows                = strings.Join(serviceTblFieldNames, ",")
	serviceTblRowsExpectAutoSet   = strings.Join(stringx.Remove(serviceTblFieldNames), ",")
	serviceTblRowsWithPlaceHolder = strings.Join(stringx.Remove(serviceTblFieldNames, "`id`"), "=?,") + "=?"
)

type (
	serviceTblModel interface {
		Insert(ctx context.Context, data *ServiceTbl) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ServiceTbl, error)
		Update(ctx context.Context, data *ServiceTbl) error
		Delete(ctx context.Context, id int64) error
	}

	defaultServiceTblModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ServiceTbl struct {
		Id        int64  `db:"id"`
		HouseId   int64  `db:"house_id"`
		Name      string `db:"name"`
		Price     int64  `db:"price"`
		Type      int64  `db:"type"`
		CreatedAt int64  `db:"created_at"`
		UpdatedAt int64  `db:"updated_at"`
		CreatedBy int64  `db:"created_by"`
		UpdatedBy int64  `db:"updated_by"`
	}
)

func newServiceTblModel(conn sqlx.SqlConn) *defaultServiceTblModel {
	return &defaultServiceTblModel{
		conn:  conn,
		table: "`service_tbl`",
	}
}

func (m *defaultServiceTblModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultServiceTblModel) FindOne(ctx context.Context, id int64) (*ServiceTbl, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", serviceTblRows, m.table)
	var resp ServiceTbl
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

func (m *defaultServiceTblModel) Insert(ctx context.Context, data *ServiceTbl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, serviceTblRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.HouseId, data.Name, data.Price, data.Type, data.CreatedAt, data.UpdatedAt, data.CreatedBy, data.UpdatedBy)
	return ret, err
}

func (m *defaultServiceTblModel) Update(ctx context.Context, data *ServiceTbl) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, serviceTblRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.HouseId, data.Name, data.Price, data.Type, data.CreatedAt, data.UpdatedAt, data.CreatedBy, data.UpdatedBy, data.Id)
	return err
}

func (m *defaultServiceTblModel) tableName() string {
	return m.table
}
