package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AlbumTblModel = (*customAlbumTblModel)(nil)

type (
	// AlbumTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAlbumTblModel.
	AlbumTblModel interface {
		albumTblModel
		withSession(session sqlx.Session) AlbumTblModel
		FindByHouseID(ctx context.Context, houseID int64) ([]*AlbumTbl, error)
		DeleteByHouseID(ctx context.Context, houseID int64) error
		FindMultiByHouseIDs(ctx context.Context, houseIDs []int64) ([]*AlbumTbl, error)
	}

	customAlbumTblModel struct {
		*defaultAlbumTblModel
	}
)

// NewAlbumTblModel returns a model for the database table.
func NewAlbumTblModel(conn sqlx.SqlConn) AlbumTblModel {
	return &customAlbumTblModel{
		defaultAlbumTblModel: newAlbumTblModel(conn),
	}
}

func (m *customAlbumTblModel) withSession(session sqlx.Session) AlbumTblModel {
	return NewAlbumTblModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customAlbumTblModel) FindByHouseID(ctx context.Context, houseID int64) ([]*AlbumTbl, error) {
	query := fmt.Sprintf("select %s from %s where `house_id` = ?", albumTblRows, m.table)
	var resp []*AlbumTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query, houseID)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *customAlbumTblModel) DeleteByHouseID(ctx context.Context, houseID int64) error {
	query := fmt.Sprintf("delete from %s where `house_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, houseID)
	return err
}

func (m *customAlbumTblModel) FindMultiByHouseIDs(ctx context.Context, houseIDs []int64) ([]*AlbumTbl, error) {
	query := fmt.Sprintf("select %s from %s where `house_id` in (", albumTblRows, m.table)
	var vals []interface{}
	var resp []*AlbumTbl
	for _, id := range houseIDs {
		query += "?,"
		vals = append(vals, id)
	}
	query = query[:len(query)-1] + ")"
	err := m.conn.QueryRowsCtx(ctx, &resp, query, vals...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
