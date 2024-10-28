package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilterContactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilterContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilterContactLogic {
	return &FilterContactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterContactLogic) FilterContact(req *types.FilterContactReq) (resp *types.FilterContactRes, err error) {
	l.Logger.Info("FilterContact: ", req)

	var userID, renterID, lessorID int64
	var contacts []types.Contact

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterContactRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	userModel, err := l.svcCtx.AccountFunction.GetUserByID(userID)
	if err != nil || userModel == nil {
		l.Logger.Error(err)
		return &types.FilterContactRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if userModel.Role.Int64 == common.USER_ROLE_LESSOR {
		lessorID = userID
	} else if userModel.Role.Int64 == common.USER_ROLE_RENTER {
		renterID = userID
	}

	total, err := l.svcCtx.ContactModel.CountByUser(l.ctx, renterID, lessorID, req.From, req.To, req.Status)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterContactRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	contactModels, err := l.svcCtx.ContactModel.FindMultiByUser(l.ctx, renterID, lessorID, req.From, req.To, req.Status, req.Limit, req.Offset)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterContactRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, contact := range contactModels {
		houseModel, err := l.svcCtx.HouseModel.FindOne(l.ctx, contact.HouseId.Int64)
		if err != nil || houseModel == nil {
			l.Logger.Error(err)
			return &types.FilterContactRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		renterModel, err := l.svcCtx.AccountFunction.GetUserByID(contact.RenterId.Int64)
		if err != nil || renterModel == nil {
			l.Logger.Error(err)
			return &types.FilterContactRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		lessorModel, err := l.svcCtx.AccountFunction.GetUserByID(contact.LessorId.Int64)
		if err != nil || lessorModel == nil {
			l.Logger.Error(err)
			return &types.FilterContactRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		contacts = append(contacts, types.Contact{
			ID:          contact.Id,
			HouseID:     contact.HouseId.Int64,
			HouseName:   houseModel.Name.String,
			RenterID:    contact.RenterId.Int64,
			RenterName:  renterModel.FullName.String,
			RenterPhone: renterModel.Phone,
			LessorID:    contact.LessorId.Int64,
			LessorName:  lessorModel.FullName.String,
			LessorPhone: lessorModel.Phone,
			Datetime:    contact.Datetime.Int64,
			Status:      contact.Status,
		})
	}

	l.Logger.Info("FilterContact Success:", userID)
	return &types.FilterContactRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Contacts: contacts,
		Total:    total,
	}, nil
}
