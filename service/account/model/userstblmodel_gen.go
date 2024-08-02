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
	usersTblFieldNames          = builder.RawFieldNames(&UsersTbl{}, true)
	usersTblRows                = strings.Join(usersTblFieldNames, ",")
	usersTblRowsExpectAutoSet   = strings.Join(stringx.Remove(usersTblFieldNames), ",")
	usersTblRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(usersTblFieldNames, "user_id"))
)

type (
	usersTblModel interface {
		Insert(ctx context.Context, data *UsersTbl) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*UsersTbl, error)
		FindOneByEmail(ctx context.Context, email string) (*UsersTbl, error)
		FindOneByUsername(ctx context.Context, username string) (*UsersTbl, error)
		Update(ctx context.Context, data *UsersTbl) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultUsersTblModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UsersTbl struct {
		UserId       int64         `db:"user_id"`
		ProfileId    sql.NullInt64 `db:"profile_id"`
		Username     string        `db:"username"`
		PasswordHash string        `db:"password_hash"`
		Email        string        `db:"email"`
		Role         sql.NullInt64 `db:"role"`
	}
)

func newUsersTblModel(conn sqlx.SqlConn) *defaultUsersTblModel {
	return &defaultUsersTblModel{
		conn:  conn,
		table: `"public"."users_tbl"`,
	}
}

func (m *defaultUsersTblModel) Delete(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("delete from %s where user_id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultUsersTblModel) FindOne(ctx context.Context, userId int64) (*UsersTbl, error) {
	query := fmt.Sprintf("select %s from %s where user_id = $1 limit 1", usersTblRows, m.table)
	var resp UsersTbl
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersTblModel) FindOneByEmail(ctx context.Context, email string) (*UsersTbl, error) {
	var resp UsersTbl
	query := fmt.Sprintf("select %s from %s where email = $1 limit 1", usersTblRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, email)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersTblModel) FindOneByUsername(ctx context.Context, username string) (*UsersTbl, error) {
	var resp UsersTbl
	query := fmt.Sprintf("select %s from %s where username = $1 limit 1", usersTblRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersTblModel) Insert(ctx context.Context, data *UsersTbl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6)", m.table, usersTblRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.ProfileId, data.Username, data.PasswordHash, data.Email, data.Role)
	return ret, err
}

func (m *defaultUsersTblModel) Update(ctx context.Context, newData *UsersTbl) error {
	query := fmt.Sprintf("update %s set %s where user_id = $1", m.table, usersTblRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.UserId, newData.ProfileId, newData.Username, newData.PasswordHash, newData.Email, newData.Role)
	return err
}

func (m *defaultUsersTblModel) tableName() string {
	return m.table
}
