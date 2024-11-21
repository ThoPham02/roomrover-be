package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
	"roomrover/service/contract/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListBillDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListBillDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListBillDetailLogic {
	return &GetListBillDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListBillDetailLogic) GetListBillDetail(req *types.GetListBillDetailReq) (resp *types.GetListBillDetailRes, err error) {
	l.Logger.Info("GetListBillDetailReq", req)

	var userID int64
	var billDetailModels []*model.BillDetailTbl
	var billDetails []types.BillDetail

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetListBillDetailRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	billDetailModels, err = l.svcCtx.BillDetailModel.GetDetailByBillID(l.ctx, req.BillID)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetListBillDetailRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, billDetailModel := range billDetailModels {
		billDetails = append(billDetails, types.BillDetail{
			BillDetailID: billDetailModel.Id,
			BillID:       billDetailModel.BillId.Int64,
			Name:         billDetailModel.Name.String,
			Price:        billDetailModel.Price.Int64,
			Type:         billDetailModel.Type.Int64,
			OldIndex:     billDetailModel.OldIndex.Int64,
			NewIndex:     billDetailModel.NewIndex.Int64,
			Imgurl:       billDetailModel.ImgUrl.String,
			Quantity:     billDetailModel.Quantity.Int64,
		})
	}

	l.Logger.Info("GetListBillDetailReq", userID)
	return &types.GetListBillDetailRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		BillDetails: billDetails,
	}, nil
}
