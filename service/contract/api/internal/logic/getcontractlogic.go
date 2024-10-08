package logic

import (
	"context"

	"roomrover/common"
	accountModel "roomrover/service/account/model"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
	"roomrover/service/contract/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContractLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get contract
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
	var contractDetails []types.ContractDetail
	var contractRenters []types.ContractRenter

	var contractModel *model.ContractTbl
	var renter *accountModel.UserTbl
	var lessor *accountModel.UserTbl
	var paymentModel *model.PaymentTbl
	var contractDetailModels []*model.ContractDetailTbl
	var contractRentersModels []*model.ContractRenterTbl

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

	renter, err = l.svcCtx.AccountFunction.GetUserByID(contractModel.RenterId.Int64)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	lessor, err = l.svcCtx.AccountFunction.GetUserByID(contractModel.LessorId.Int64)
	if err != nil {
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

	contractDetailModels, err = l.svcCtx.ContractDetailModel.GetContractDetailByContractID(l.ctx, contractModel.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, contractDetailModel := range contractDetailModels {
		contractDetail := types.ContractDetail{
			ID:         contractDetailModel.Id,
			ContractID: contractDetailModel.ContractId.Int64,
			Name:       contractDetailModel.Name.String,
			Price:      contractDetailModel.Price.Int64,
			Type:       contractDetailModel.Type.Int64,
		}
		contractDetails = append(contractDetails, contractDetail)
	}

	contractRentersModels, err = l.svcCtx.ContractRenterModel.GetRenterByContractID(l.ctx, contractModel.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, contractRenterModel := range contractRentersModels {
		userModel, err := l.svcCtx.AccountFunction.GetUserByID(contractRenterModel.UserId.Int64)
		if err != nil {
			l.Logger.Error(err)
			return &types.GetContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		contractRenter := types.ContractRenter{
			ID:         contractRenterModel.Id,
			ContractID: contractRenterModel.ContractId.Int64,
			RenterID:   userModel.Id,
			Name:       userModel.FullName.String,
			Phone:      userModel.Phone,
		}
		contractRenters = append(contractRenters, contractRenter)
	}

	contract = types.Contract{
		ContractID:    contractModel.Id,
		Code:          contractModel.Code.String,
		Status:        contractModel.Status.Int64,
		RenterID:      contractModel.RenterId.Int64,
		RenterPhone:   renter.Phone,
		RenterNumber:  contractModel.RenterNumber.String,
		RenterDate:    contractModel.RenterDate.Int64,
		RenterAddress: contractModel.RenterAddress.String,
		RenterName:    contractModel.RenterName.String,
		LessorID:      contractModel.LessorId.Int64,
		LessorPhone:   lessor.Phone,
		LessorNumber:  contractModel.LessorNumber.String,
		LessorDate:    contractModel.LessorDate.Int64,
		LessorAddress: contractModel.LessorAddress.String,
		LessorName:    contractModel.LessorName.String,
		// Room: types.Room{
		// 	RoomID:   0,
		// 	Name:     "",
		// 	Status:   0,
		// 	Capacity: 0,
		// 	EIndex:   0,
		// 	WIndex:   0,
		// },
		CheckIn:         contractModel.CheckIn.Int64,
		Duration:        contractModel.Duration.Int64,
		Purpose:         contractModel.Purpose.String,
		ContractRenters: contractRenters,
		ContractDetails: contractDetails,
		Payment: types.Payment{
			PaymentID:   paymentModel.Id,
			ContractID:  paymentModel.ContractId.Int64,
			Amount:      paymentModel.Amount.Int64,
			Discount:    paymentModel.Discount.Int64,
			Deposit:     paymentModel.Deposit.Int64,
			DepositDate: paymentModel.DepositDate.Int64,
			NextBill:    paymentModel.NextBill.Int64,
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
