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
	var payment types.Payment
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
		Id:            l.svcCtx.ObjSync.GenServiceObjID(),
		Code:          sql.NullString{},
		Status:        sql.NullInt64{Valid: true, Int64: common.CONTRACT_STATUS_PENDING},
		RenterId:      sql.NullInt64{},
		RenterNumber:  sql.NullString{},
		RenterDate:    sql.NullInt64{},
		RenterAddress: sql.NullString{},
		RenterName:    sql.NullString{},
		LessorId:      sql.NullInt64{},
		LessorNumber:  sql.NullString{},
		LessorDate:    sql.NullInt64{},
		LessorAddress: sql.NullString{},
		LessorName:    sql.NullString{},
		RoomId:        sql.NullInt64{Valid: true, Int64: req.RoomID},
		CheckIn:       sql.NullInt64{},
		Duration:      sql.NullInt64{},
		Purpose:       sql.NullString{},
		CreatedAt:     sql.NullInt64{Valid: true, Int64: currentTime},
		UpdatedAt:     sql.NullInt64{Valid: true, Int64: currentTime},
		CreatedBy:     sql.NullInt64{Valid: true, Int64: userID},
		UpdatedBy:     sql.NullInt64{Valid: true, Int64: userID},
	}

	for _, renter := range contractRenters {
		contractRenterModels = append(contractRenterModels, &model.ContractRenterTbl{
			Id:         l.svcCtx.ObjSync.GenServiceObjID(),
			ContractId: sql.NullInt64{Valid: true, Int64: contractModel.Id},
			UserId:     sql.NullInt64{Valid: true, Int64: renter.RenterID},
		})
	}
	for _, detail := range contractDetails {
		contractDetailModels = append(contractDetailModels, &model.ContractDetailTbl{
			Id:         l.svcCtx.ObjSync.GenServiceObjID(),
			ContractId: sql.NullInt64{Valid: true, Int64: contractModel.Id},
			Name:       sql.NullString{Valid: true, String: detail.Name},
			Type:       sql.NullInt64{Valid: true, Int64: detail.Type},
			Price:      sql.NullInt64{Valid: true, Int64: detail.Price},
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

	payment = types.Payment{

	}

	contract = types.Contract{
		ContractID:      contractModel.Id,
		Code:            contractModel.Code.String,
		Status:          contractModel.Status.Int64,
		RenterID:        contractModel.RenterId.Int64,
		RenterNumber:    contractModel.RenterNumber.String,
		RenterDate:      contractModel.RenterDate.Int64,
		RenterAddress:   contractModel.RenterAddress.String,
		RenterName:      contractModel.RenterName.String,
		LessorID:        contractModel.LessorId.Int64,
		LessorNumber:    contractModel.LessorNumber.String,
		LessorDate:      contractModel.LessorDate.Int64,
		LessorAddress:   contractModel.LessorAddress.String,
		LessorName:      contractModel.LessorName.String,
		RoomID:          contractModel.RoomId.Int64,
		CheckIn:         contractModel.CheckIn.Int64,
		Duration:        contractModel.Duration.Int64,
		Purpose:         contractModel.Purpose.String,
		ContractRenters: contractRenters,
		ContractDetails: contractDetails,
		Payment:         payment,
		CreatedAt:       contractModel.CreatedAt.Int64,
		UpdatedAt:       contractModel.UpdatedAt.Int64,
		CreatedBy:       contractModel.CreatedBy.Int64,
		UpdatedBy:       contractModel.UpdatedBy.Int64,
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
