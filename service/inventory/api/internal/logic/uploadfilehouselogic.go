package logic

import (
	"context"
	"net/http"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileHouseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

// Upload file house
func NewUploadFileHouseLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadFileHouseLogic {
	return &UploadFileHouseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadFileHouseLogic) UploadFileHouse(req *types.UploadFileHouseReq) (resp *types.UploadFileHouseRes, err error) {
	l.Logger.Info("UploadFileHouse", req)

	var userID int64

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.UploadFileHouseRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	// get file from request
	file, _, err := l.r.FormFile("file")
	if err != nil {
		l.Logger.Error(err)
		return &types.UploadFileHouseRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}
	defer file.Close()

	// upload file to cloud
	url, err := l.svcCtx.CldClient.UploadImage(l.ctx, file, userID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UploadFileHouseRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("UploadFileHouse Success", userID)
	return &types.UploadFileHouseRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Url: url,
	}, nil
}
