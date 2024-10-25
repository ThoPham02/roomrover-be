package logic

import (
	"context"

	"roomrover/common"
	accountModel "roomrover/service/account/model"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
	"roomrover/service/contract/model"
	inventoryModel "roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContractLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContractLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContractLogic {
	return &GetContractLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContractLogic) GetContract(req *types.GetContractReq) (resp *types.GetContractRes, err error) {
	l.Logger.Info("GetContract", req)

	var userID int64

	var contract types.Contract
	var paymentDetails []types.PaymentDetail
	var paymentRenters []types.PaymentRenter

	var contractModel *model.ContractTbl
	var renterModel *accountModel.UserTbl
	var lessorModel *accountModel.UserTbl
	var houseRoomModel *inventoryModel.HouseRoomTbl
	var paymentModel *model.PaymentTbl
	var paymentDetailModels []*model.PaymentDetailTbl
	var paymentRentersModels []*model.PaymentRenterTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetContractRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}
	contractModel, err = l.svcCtx.ContractModel.FindOne(l.ctx, req.ID)
	if err != nil {
		l.Logger.Error(err)
		if err == model.ErrNotFound {
			return &types.GetContractRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		return &types.GetContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	renterModel, err = l.svcCtx.AccountFunction.GetUserByID(contractModel.RenterId.Int64)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	lessorModel, err = l.svcCtx.AccountFunction.GetUserByID(contractModel.LessorId.Int64)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	houseRoomModel, err = l.svcCtx.InventFunction.GetHouseRoomByRoomID(contractModel.RoomId.Int64)
	if err != nil || houseRoomModel == nil {
		l.Logger.Error(err)
		return &types.GetContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	paymentModel, err = l.svcCtx.PaymentModel.FindByContractID(l.ctx, contractModel.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	paymentDetailModels, err = l.svcCtx.PaymentDetailModel.GetPaymentDetailByPaymentID(l.ctx, paymentModel.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, detail := range paymentDetailModels {
		paymentDetail := types.PaymentDetail{
			ID:        detail.Id,
			PaymentID: detail.PaymentId.Int64,
			Name:      detail.Name.String,
			Price:     detail.Price.Int64,
			Type:      detail.Type.Int64,
		}
		paymentDetails = append(paymentDetails, paymentDetail)
	}

	paymentRentersModels, err = l.svcCtx.PaymentRenterModel.GetRenterByPaymentID(l.ctx, paymentModel.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, renter := range paymentRentersModels {
		userModel, err := l.svcCtx.AccountFunction.GetUserByID(renter.UserId.Int64)
		if err != nil {
			l.Logger.Error(err)
			return &types.GetContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		paymentRenters = append(paymentRenters, types.PaymentRenter{
			ID:        renter.Id,
			PaymentID: renter.PaymentId.Int64,
			RenterID:  userModel.Id,
			Name:      userModel.FullName.String,
			Phone:     userModel.Phone,
		})
	}

	contract = types.Contract{
		ContractID: contractModel.Id,
		Code:       contractModel.Code.String,
		Status:     contractModel.Status.Int64,
		Renter: types.User{
			UserID:      renterModel.Id,
			Phone:       renterModel.Phone,
			Role:        renterModel.Role.Int64,
			Status:      renterModel.Status,
			Address:     renterModel.Address.String,
			FullName:    renterModel.FullName.String,
			AvatarUrl:   renterModel.AvatarUrl.String,
			Birthday:    renterModel.Birthday.Int64,
			Gender:      renterModel.Gender.Int64,
			CccdNumber:  contractModel.RenterNumber.String,
			CccdDate:    contractModel.RenterDate.Int64,
			CccdAddress: contractModel.RenterAddress.String,
			CreatedAt:   renterModel.CreatedAt.Int64,
			UpdatedAt:   renterModel.UpdatedAt.Int64,
		},
		Lessor: types.User{
			UserID:      lessorModel.Id,
			Phone:       lessorModel.Phone,
			Role:        lessorModel.Role.Int64,
			Status:      lessorModel.Status,
			Address:     lessorModel.Address.String,
			FullName:    lessorModel.FullName.String,
			AvatarUrl:   lessorModel.AvatarUrl.String,
			Birthday:    lessorModel.Birthday.Int64,
			Gender:      lessorModel.Gender.Int64,
			CccdNumber:  contractModel.LessorNumber.String,
			CccdDate:    contractModel.LessorDate.Int64,
			CccdAddress: contractModel.LessorAddress.String,
			CreatedAt:   lessorModel.CreatedAt.Int64,
			UpdatedAt:   lessorModel.UpdatedAt.Int64,
		},
		Room: types.Room{
			RoomID:     houseRoomModel.Id,
			Name:       houseRoomModel.HouseRoomName.String,
			ProvinceID: houseRoomModel.ProvinceID.Int64,
			DistrictID: houseRoomModel.DistrictID.Int64,
			WardID:     houseRoomModel.WardID.Int64,
			Address:    houseRoomModel.Address.String,
			Area:       houseRoomModel.Area.Int64,
			Price:      houseRoomModel.Price.Int64,
			Type:       houseRoomModel.Type.Int64,
			EIndex:     houseRoomModel.EIndex.Int64,
			WIndex:     houseRoomModel.WIndex.Int64,
		},
		CheckIn:  contractModel.CheckIn.Int64,
		Duration: contractModel.Duration.Int64,
		Purpose:  contractModel.Purpose.String,
		Payment: types.Payment{
			PaymentID:      paymentModel.Id,
			ContractID:     paymentModel.ContractId,
			Amount:         paymentModel.Amount,
			Discount:       paymentModel.Discount,
			Deposit:        paymentModel.Deposit,
			DepositDate:    paymentModel.DepositDate,
			NextBill:       paymentModel.NextBill,
			PaymentRenters: paymentRenters,
			PaymentDetails: paymentDetails,
		},
		CreatedAt: contractModel.CreatedAt.Int64,
		UpdatedAt: contractModel.UpdatedAt.Int64,
		CreatedBy: contractModel.CreatedBy.Int64,
		UpdatedBy: contractModel.UpdatedBy.Int64,
	}

	l.Logger.Info("GetContract Success: ", userID)
	return &types.GetContractRes{
		Contract: contract,
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
