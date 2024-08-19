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
