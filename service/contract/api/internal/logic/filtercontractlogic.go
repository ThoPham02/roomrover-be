package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
	"roomrover/service/contract/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilterContractLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilterContractLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilterContractLogic {
	return &FilterContractLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilterContractLogic) FilterContract(req *types.FilterContractReq) (resp *types.FilterContractRes, err error) {
	l.Logger.Info("FilterContract: ", req)

	var userID int64
	var count int64
	var listUserID []int64
	var mapUserPhone = map[int64]string{}

	var listContract []types.Contract

	var contractModels []*model.ContractTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterContractRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	count, err = l.svcCtx.ContractModel.CountContractByCondition(l.ctx, userID, 0, req.Search, req.Status, req.CreateFrom, req.CreateTo)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if count == 0 {
		return &types.FilterContractRes{
			Result: types.Result{
				Code:    common.SUCCESS_CODE,
				Message: common.SUCCESS_MESS,
			},
		}, nil
	}

	contractModels, err = l.svcCtx.ContractModel.FindContractByCondition(l.ctx, userID, 0, req.Search, req.Status, req.CreateFrom, req.CreateTo, req.Offset, req.Limit)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, contractModel := range contractModels {
		listUserID = append(listUserID, contractModel.RenterId.Int64)
		listUserID = append(listUserID, contractModel.LessorId.Int64)
	}

	userModels, err := l.svcCtx.AccountFunction.GetUsersByIDs(listUserID)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, userModel := range userModels {
		mapUserPhone[userModel.Id] = userModel.Phone
	}

	for _, contractModel := range contractModels {
		paymentModel, err := l.svcCtx.PaymentModel.FindByContractID(l.ctx, contractModel.Id)
		if err != nil {
			l.Logger.Error(err)
			return &types.FilterContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		listContract = append(listContract, types.Contract{
			ContractID:    contractModel.Id,
			Code:          contractModel.Code.String,
			Status:        contractModel.Status.Int64,
			RenterID:      contractModel.RenterId.Int64,
			RenterPhone:   mapUserPhone[contractModel.RenterId.Int64],
			RenterNumber:  contractModel.RenterNumber.String,
			RenterDate:    contractModel.RenterDate.Int64,
			RenterAddress: contractModel.RenterAddress.String,
			RenterName:    contractModel.RenterName.String,
			LessorID:      contractModel.LessorId.Int64,
			LessorPhone:   mapUserPhone[contractModel.LessorId.Int64],
			LessorNumber:  contractModel.LessorNumber.String,
			LessorDate:    contractModel.LessorDate.Int64,
			LessorAddress: contractModel.LessorAddress.String,
			LessorName:    contractModel.LessorName.String,
			Room: types.Room{
				RoomID:   0,
				Name:     "",
				Status:   0,
				Capacity: 0,
				EIndex:   0,
				WIndex:   0,
			},
			CheckIn:  contractModel.CheckIn.Int64,
			Duration: contractModel.Duration.Int64,
			Purpose:  contractModel.Purpose.String,
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
		})
	}

	l.Logger.Info("FilterContract Success:", userID)
	return &types.FilterContractRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Contracts: listContract,
		Total:     count,
	}, nil
}
