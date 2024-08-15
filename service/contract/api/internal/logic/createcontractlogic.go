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

type CreateContractLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateContractLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContractLogic {
	return &CreateContractLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateContractLogic) CreateContract(req *types.CreateContractReq) (resp *types.CreateContractRes, err error) {
	l.Logger.Info("CreateContract", req)

	var userID int64
	var currentTime int64

	var contract types.Contract
	var contractRenters []types.ContractRenter
	var contractDetails []types.ContractDetail

	var contractModel *model.ContractTbl
	var contractDetailModels []*model.ContractDetailTbl
	var contractRenterModels []*model.ContractRenterTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	_, err = l.svcCtx.InventFunction.GetRoomByID(req.RoomID)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.CreateContractRes{
				Result: types.Result{
					Code:    common.ROOM_NOT_FOUND_CODE,
					Message: common.ROOM_NOT_FOUND_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	err = json.Unmarshal([]byte(req.ContractRenter), &contractRenters)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}
	err = json.Unmarshal([]byte(req.ContractDetail), &contractDetails)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	contractModel = &model.ContractTbl{
		Id:          l.svcCtx.ObjSync.GenServiceObjID(),
		RoomId:      req.RoomID,
		Status:      common.CONTRACT_STATUS_PENDING,
		ContractUrl: sql.NullString{Valid: true, String: req.ContractUrl},
		Description: req.Description,
		Start:       req.Start,
		End:         req.End,
		NextBill:    common.GetNextMonthDate(req.Start, 1),
		Type:        req.Type,
		Deposit:     req.Deposit,
		Deadline:    req.Deadline,
		DepositUrl:  sql.NullString{Valid: true, String: req.DepositUrl},
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
		CreatedBy:   userID,
		UpdatedBy:   userID,
	}

	for _, renter := range contractRenters {
		contractRenterModels = append(contractRenterModels, &model.ContractRenterTbl{
			Id:         l.svcCtx.ObjSync.GenServiceObjID(),
			ContractId: contractModel.Id,
			RenterId:   renter.RenterID,
			Type:       renter.Type,
		})
	}
	for _, detail := range contractDetails {
		contractDetailModels = append(contractDetailModels, &model.ContractDetailTbl{
			Id:        l.svcCtx.ObjSync.GenServiceObjID(),
			ContractId: contractModel.Id,
			ServiceId: detail.ServiceID,
			Price:     detail.Price,
			Type:      detail.Type,
			Index:     detail.Index,
		})
	}

	err = l.svcCtx.ContractDetailModel.InsertMulti(l.ctx, contractDetailModels)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	err = l.svcCtx.ContractRenterModel.InsertMulti(l.ctx, contractRenterModels)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	_, err = l.svcCtx.ContractModel.Insert(l.ctx, contractModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	contract = types.Contract{
		ContractID:      contractModel.Id,
		RoomID:          contractModel.RoomId,
		Description:     contractModel.Description,
		ContractUrl:     contractModel.ContractUrl.String,
		Start:           contractModel.Start,
		End:             contractModel.End,
		Status:          contractModel.Status,
		Type:            contractModel.Type,
		Deposit:         contractModel.Deposit,
		Deadline:        contractModel.Deadline,
		CreatedAt:       contractModel.CreatedAt,
		UpdatedAt:       contractModel.UpdatedAt,
		CreatedBy:       contractModel.CreatedBy,
		UpdatedBy:       contractModel.UpdatedBy,
		ContractRenters: contractRenters,
		ContractDetails: contractDetails,
	}

	l.Logger.Info("CreateContract Success", userID)
	return &types.CreateContractRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Contract: contract,
	}, nil
}
