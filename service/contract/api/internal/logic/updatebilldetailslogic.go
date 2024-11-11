package logic

import (
	"context"
	"database/sql"
	"encoding/json"

	"roomrover/common"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
	"roomrover/service/contract/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBillDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBillDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBillDetailsLogic {
	return &UpdateBillDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBillDetailsLogic) UpdateBillDetails(req *types.UpdateBillDetailsReq) (resp *types.UpdateBillDetailsRes, err error) {
	l.Logger.Info("UpdateBillDetails", req)

	var userID int64
	var billDetailModels []*model.BillDetailTbl
	var billDetails []types.BillDetail
	var billModel *model.BillTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateBillDetailsRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	err = json.Unmarshal([]byte(req.BillDetails), &billDetails)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateBillDetailsRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	for _, billDetail := range billDetails {
		var billDetailModel *model.BillDetailTbl

		billDetailModel, err = l.svcCtx.BillDetailModel.FindOne(l.ctx, billDetail.BillDetailID)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateBillDetailsRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		billDetailModel.Quantity = sql.NullInt64{Valid: true, Int64: billDetail.NewIndex - billDetail.OldIndex}
		billDetailModel.NewIndex = sql.NullInt64{Valid: true, Int64: billDetail.NewIndex}
		billDetailModel.ImgUrl = sql.NullString{Valid: true, String: billDetail.Imgurl}

		billDetailModels = append(billDetailModels, billDetailModel)
	}

	for _, billDetailModel := range billDetailModels {
		err = l.svcCtx.BillDetailModel.Update(l.ctx, billDetailModel)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateBillDetailsRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}

	billModel, err = l.svcCtx.BillModel.FindOne(l.ctx, req.BillID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateBillDetailsRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	
	billDetailModels, err = l.svcCtx.BillDetailModel.GetDetailByBillID(l.ctx, req.BillID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateBillDetailsRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, billDetailModel := range billDetailModels {
	}

	l.Logger.Info("UpdateBillDetails Success: ", userID)
	return &types.UpdateBillDetailsRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
