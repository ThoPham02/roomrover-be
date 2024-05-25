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
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`user_id`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), "=?,") + "=?"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*Users, error)
		FindOneByEmail(ctx context.Context, email string) (*Users, error)
		FindOneByUsername(ctx context.Context, username string) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultUsersModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Users struct {
		UserId       int64         `db:"user_id"`
		ProfileId    sql.NullInt64 `db:"profile_id"`
		Username     string        `db:"username"`
		PasswordHash string        `db:"password_hash"`
		Email        string        `db:"email"`
		Role         sql.NullInt64 `db:"role"`
	}
)

func newUsersModel(conn sqlx.SqlConn) *defaultUsersModel {
	return &defaultUsersModel{
		conn:  conn,
		table: "`users`",
	}
}

func (m *defaultUsersModel) Delete(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, userId int64) (*Users, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", usersRows, m.table)
	var resp Users
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByEmail(ctx context.Context, email string) (*Users, error) {
	var resp Users
	query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", usersRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, email)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByUsername(ctx context.Context, username string) (*Users, error) {
	var resp Users
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", usersRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.ProfileId, data.Username, data.PasswordHash, data.Email, data.Role)
	return ret, err
}

func (m *defaultUsersModel) Update(ctx context.Context, newData *Users) error {
	query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, usersRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.ProfileId, newData.Username, newData.PasswordHash, newData.Email, newData.Role, newData.UserId)
	return err
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}
