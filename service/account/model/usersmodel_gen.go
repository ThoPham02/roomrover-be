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
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`user_id`"), "=?,") + "=?"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*Users, error)
		FindOneByPhone(ctx context.Context, phone string) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultUsersModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Users struct {
		UserId       int64          `db:"user_id"`
		Phone        string         `db:"phone"`
		PasswordHash string         `db:"password_hash"`
		Role         sql.NullInt64  `db:"role"`
		Status       int64          `db:"status"`
		Address      sql.NullString `db:"address"`
		FullName     sql.NullString `db:"full_name"`
		AvatarUrl    sql.NullString `db:"avatar_url"`
		Birthday     sql.NullInt64  `db:"birthday"`
		Gender       sql.NullInt64  `db:"gender"`
		CreatedAt    int64          `db:"created_at"`
		UpdatedAt    int64          `db:"updated_at"`
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

func (m *defaultUsersModel) FindOneByPhone(ctx context.Context, phone string) (*Users, error) {
	var resp Users
	query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", usersRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, phone)
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
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Phone, data.PasswordHash, data.Role, data.Status, data.Address, data.FullName, data.AvatarUrl, data.Birthday, data.Gender, data.CreatedAt, data.UpdatedAt)
	return ret, err
}

func (m *defaultUsersModel) Update(ctx context.Context, newData *Users) error {
	query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, usersRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Phone, newData.PasswordHash, newData.Role, newData.Status, newData.Address, newData.FullName, newData.AvatarUrl, newData.Birthday, newData.Gender, newData.CreatedAt, newData.UpdatedAt, newData.UserId)
	return err
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}
