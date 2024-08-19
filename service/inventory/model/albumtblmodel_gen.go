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
	albumTblFieldNames          = builder.RawFieldNames(&AlbumTbl{})
	albumTblRows                = strings.Join(albumTblFieldNames, ",")
	albumTblRowsExpectAutoSet   = strings.Join(stringx.Remove(albumTblFieldNames), ",")
	albumTblRowsWithPlaceHolder = strings.Join(stringx.Remove(albumTblFieldNames, "`id`"), "=?,") + "=?"
)

type (
	albumTblModel interface {
		Insert(ctx context.Context, data *AlbumTbl) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AlbumTbl, error)
		Update(ctx context.Context, data *AlbumTbl) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAlbumTblModel struct {
		conn  sqlx.SqlConn
		table string
	}

	AlbumTbl struct {
		Id        int64  `db:"id"`
		HouseId   int64  `db:"house_id"`
		Url       string `db:"url"`
		Type      int64  `db:"type"`
		CreatedAt int64  `db:"created_at"`
		UpdatedAt int64  `db:"updated_at"`
		CreatedBy int64  `db:"created_by"`
		UpdatedBy int64  `db:"updated_by"`
	}
)

func newAlbumTblModel(conn sqlx.SqlConn) *defaultAlbumTblModel {
	return &defaultAlbumTblModel{
		conn:  conn,
		table: "`album_tbl`",
	}
}

func (m *defaultAlbumTblModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAlbumTblModel) FindOne(ctx context.Context, id int64) (*AlbumTbl, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", albumTblRows, m.table)
	var resp AlbumTbl
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

func (m *defaultAlbumTblModel) Insert(ctx context.Context, data *AlbumTbl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, albumTblRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.HouseId, data.Url, data.Type, data.CreatedAt, data.UpdatedAt, data.CreatedBy, data.UpdatedBy)
	return ret, err
}

func (m *defaultAlbumTblModel) Update(ctx context.Context, data *AlbumTbl) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, albumTblRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.HouseId, data.Url, data.Type, data.CreatedAt, data.UpdatedAt, data.CreatedBy, data.UpdatedBy, data.Id)
	return err
}

func (m *defaultAlbumTblModel) tableName() string {
	return m.table
}
