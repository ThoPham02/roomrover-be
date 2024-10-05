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
	contractRenterTblFieldNames          = builder.RawFieldNames(&ContractRenterTbl{})
	contractRenterTblRows                = strings.Join(contractRenterTblFieldNames, ",")
	contractRenterTblRowsExpectAutoSet   = strings.Join(stringx.Remove(contractRenterTblFieldNames), ",")
	contractRenterTblRowsWithPlaceHolder = strings.Join(stringx.Remove(contractRenterTblFieldNames, "`id`"), "=?,") + "=?"
)

type (
	contractRenterTblModel interface {
		Insert(ctx context.Context, data *ContractRenterTbl) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ContractRenterTbl, error)
		Update(ctx context.Context, data *ContractRenterTbl) error
		Delete(ctx context.Context, id int64) error
	}

	defaultContractRenterTblModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ContractRenterTbl struct {
		Id         int64         `db:"id"`
		ContractId sql.NullInt64 `db:"contract_id"`
		UserId     sql.NullInt64 `db:"user_id"`
	}
)

func newContractRenterTblModel(conn sqlx.SqlConn) *defaultContractRenterTblModel {
	return &defaultContractRenterTblModel{
		conn:  conn,
		table: "`contract_renter_tbl`",
	}
}

func (m *defaultContractRenterTblModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultContractRenterTblModel) FindOne(ctx context.Context, id int64) (*ContractRenterTbl, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", contractRenterTblRows, m.table)
	var resp ContractRenterTbl
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

func (m *defaultContractRenterTblModel) Insert(ctx context.Context, data *ContractRenterTbl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, contractRenterTblRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.ContractId, data.UserId)
	return ret, err
}

func (m *defaultContractRenterTblModel) Update(ctx context.Context, data *ContractRenterTbl) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, contractRenterTblRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ContractId, data.UserId, data.Id)
	return err
}

func (m *defaultContractRenterTblModel) tableName() string {
	return m.table
}
