package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListRenterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListRenterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListRenterLogic {
	return &GetListRenterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetListRenterLogic) GetListRenter(req *types.GetListRenterReq) (resp *types.GetListRenterRes, err error) {
	l.Logger.Info("GetListRenter ", req)

	var userID int64
	var renters []types.RenterContact
	var total int

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetListRenterRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	total, err = l.svcCtx.PaymentRenterModel.CountRenterContacts(l.ctx, userID, req.Search)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetListRenterRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if total == 0 {
		return &types.GetListRenterRes{
			Result: types.Result{
				Code:    common.SUCCESS_CODE,
				Message: common.SUCCESS_MESS,
			},
		}, nil
	}

	renterContactModels, err := l.svcCtx.PaymentRenterModel.FilterRenterContacts(l.ctx, userID, req.Search, req.Limit, req.Offset)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetListRenterRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, renterContactModel := range renterContactModels {
		renter := types.RenterContact{
			ID:          renterContactModel.Id,
			RoomName:    renterContactModel.RoomName,
			Name:        renterContactModel.Name,
			Phone:       renterContactModel.Phone,
			CccdNumber:  renterContactModel.CccdNumber,
			CccdDate:    renterContactModel.CccdDate,
			CccdAddress: renterContactModel.CccdAddress,
			Status:      renterContactModel.Status,
		}
		renters = append(renters, renter)
	}

	l.Logger.Info("GetListRenter Success", userID)
	return &types.GetListRenterRes{
		Result:  types.Result{Code: common.SUCCESS_CODE, Message: common.SUCCESS_MESS},
		Renters: renters,
		Total:   total,
	}, nil
}
