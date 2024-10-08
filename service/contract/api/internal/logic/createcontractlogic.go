package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"roomrover/common"
	accountModel "roomrover/service/account/model"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
	"roomrover/service/contract/model"
	inventoryModel "roomrover/service/inventory/model"

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
	l.Logger.Info("CreateContractLogic: ", req)

	var userID int64
	var currentTime = common.GetCurrentTime()
	var contractCode string
	var codeTime = time.Now().Format("20060102")

	var contract types.Contract
	var contractRenters []types.ContractRenter
	var contractDetails []types.ContractDetail

	var renterModel *accountModel.UserTbl
	var lessorModel *accountModel.UserTbl
	var roomModel *inventoryModel.RoomTbl
	// var houseModel *inventoryModel.HouseTbl
	var serviceModels []*inventoryModel.ServiceTbl
	var contractModel *model.ContractTbl
	var paymentModel *model.PaymentTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
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
	for _, renter := range contractRenters {
		userCheck, err := l.svcCtx.AccountFunction.GetUserByID(renter.RenterID)
		if err != nil {
			l.Logger.Error(err)
			return &types.CreateContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
		if userCheck == nil {
			return &types.CreateContractRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
	}

	renterModel, err = l.svcCtx.AccountFunction.GetUserByID(req.RenterID)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	lessorModel, err = l.svcCtx.AccountFunction.GetUserByID(req.LessorID)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	roomModel, err = l.svcCtx.InventFunction.GetRoomByID(req.RoomID)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if renterModel == nil || lessorModel == nil || roomModel == nil {
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	serviceModels, err = l.svcCtx.InventFunction.GetSericesByRoom(req.RoomID)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	renterModel.CCCDNumber = sql.NullString{String: req.RenterNumber, Valid: true}
	renterModel.CCCDDate = sql.NullInt64{Int64: req.RenterDate, Valid: true}
	renterModel.CCCDAddress = sql.NullString{String: req.RenterAddress, Valid: true}
	renterModel.FullName = sql.NullString{String: req.RenterName, Valid: true}
	renterModel.UpdatedAt = sql.NullInt64{Int64: currentTime, Valid: true}
	err = l.svcCtx.AccountFunction.UpdateUser(renterModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	lessorModel.CCCDNumber = sql.NullString{String: req.LessorNumber, Valid: true}
	lessorModel.CCCDDate = sql.NullInt64{Int64: req.LessorDate, Valid: true}
	lessorModel.CCCDAddress = sql.NullString{String: req.LessorAddress, Valid: true}
	lessorModel.FullName = sql.NullString{String: req.LessorName, Valid: true}
	lessorModel.UpdatedAt = sql.NullInt64{Int64: currentTime, Valid: true}
	err = l.svcCtx.AccountFunction.UpdateUser(lessorModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	keyContractCode := "CONTRACT|" + strconv.FormatInt(userID, 10) + "|" + codeTime + "|"
	count, err := l.svcCtx.ContractRedis.IncreaseContractCode(l.ctx, keyContractCode, 24*time.Hour)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	contractCode = codeTime + fmt.Sprintf("%03d", int64(count))

	contractModel = &model.ContractTbl{
		Id:            l.svcCtx.ObjSync.GenServiceObjID(),
		Code:          sql.NullString{Valid: true, String: contractCode},
		Status:        sql.NullInt64{Valid: true, Int64: common.CONTRACT_STATUS_DRAF},
		RenterId:      sql.NullInt64{Valid: true, Int64: renterModel.Id},
		RenterNumber:  sql.NullString{Valid: true, String: renterModel.CCCDNumber.String},
		RenterDate:    sql.NullInt64{Valid: true, Int64: renterModel.CCCDDate.Int64},
		RenterAddress: sql.NullString{Valid: true, String: renterModel.Address.String},
		RenterName:    sql.NullString{Valid: true, String: renterModel.FullName.String},
		LessorId:      sql.NullInt64{Valid: true, Int64: lessorModel.Id},
		LessorNumber:  sql.NullString{Valid: true, String: lessorModel.CCCDNumber.String},
		LessorDate:    sql.NullInt64{Valid: true, Int64: lessorModel.CCCDDate.Int64},
		LessorAddress: sql.NullString{Valid: true, String: lessorModel.Address.String},
		LessorName:    sql.NullString{Valid: true, String: lessorModel.FullName.String},
		RoomId:        sql.NullInt64{Valid: true, Int64: roomModel.Id},
		CheckIn:       sql.NullInt64{Valid: true, Int64: req.CheckIn},
		Duration:      sql.NullInt64{Valid: true, Int64: req.Duration},
		Purpose:       sql.NullString{Valid: true, String: req.Purpose},
		CreatedAt:     sql.NullInt64{Valid: true, Int64: currentTime},
		UpdatedAt:     sql.NullInt64{Valid: true, Int64: currentTime},
		CreatedBy:     sql.NullInt64{Valid: true, Int64: userID},
		UpdatedBy:     sql.NullInt64{Valid: true, Int64: userID},
	}

	paymentModel = &model.PaymentTbl{
		Id:          l.svcCtx.ObjSync.GenServiceObjID(),
		ContractId:  sql.NullInt64{Valid: true, Int64: contractModel.Id},
		Amount:      sql.NullInt64{Valid: true, Int64: req.Amount},
		Discount:    sql.NullInt64{Valid: true, Int64: req.Discount},
		Deposit:     sql.NullInt64{Valid: true, Int64: req.Deposit},
		DepositDate: sql.NullInt64{Valid: true, Int64: req.DepositDate},
		NextBill:    sql.NullInt64{Valid: true, Int64: common.GetNextMonthDate(req.CheckIn, 1)},
	}

	roomModel.EIndex = sql.NullInt64{Int64: req.EIndex, Valid: true}
	roomModel.WIndex = sql.NullInt64{Int64: req.WIndex, Valid: true}
	roomModel.Status = common.ROOM_STATUS_RENTED
	err = l.svcCtx.InventFunction.UpdateRoom(roomModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.CreateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	for _, service := range serviceModels {
		contractDetails = append(contractDetails, types.ContractDetail{
			ID:         userID,
			ContractID: contractModel.Id,
			Name:       service.Name.String,
			Price:      service.Price.Int64,
			Type:       service.Unit.Int64,
		})

		contractDetail := &model.ContractDetailTbl{
			Id:         l.svcCtx.ObjSync.GenServiceObjID(),
			ContractId: sql.NullInt64{Valid: true, Int64: contractModel.Id},
			Name:       sql.NullString{Valid: true, String: service.Name.String},
			Type:       sql.NullInt64{Valid: true, Int64: service.Unit.Int64},
			Price:      sql.NullInt64{Valid: true, Int64: service.Price.Int64},
		}
		_, err = l.svcCtx.ContractDetailModel.Insert(l.ctx, contractDetail)
		if err != nil {
			l.Logger.Error(err)
			return &types.CreateContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}

	for _, renter := range contractRenters {
		contractRenterModel := &model.ContractRenterTbl{
			Id:         l.svcCtx.ObjSync.GenServiceObjID(),
			ContractId: sql.NullInt64{Valid: true, Int64: contractModel.Id},
			UserId:     sql.NullInt64{Valid: true, Int64: renter.RenterID},
		}
		_, err = l.svcCtx.ContractRenterModel.Insert(l.ctx, contractRenterModel)
		if err != nil {
			l.Logger.Error(err)
			return &types.CreateContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
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
	_, err = l.svcCtx.PaymentModel.Insert(l.ctx, paymentModel)
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
		ContractID:    contractModel.Id,
		Code:          contractModel.Code.String,
		Status:        contractModel.Status.Int64,
		RenterID:      contractModel.RenterId.Int64,
		RenterPhone:   renterModel.Phone,
		RenterNumber:  contractModel.RenterNumber.String,
		RenterDate:    contractModel.RenterDate.Int64,
		RenterAddress: contractModel.RenterAddress.String,
		RenterName:    contractModel.RenterName.String,
		LessorID:      contractModel.LessorId.Int64,
		LessorPhone:   lessorModel.Phone,
		LessorNumber:  contractModel.LessorNumber.String,
		LessorDate:    contractModel.LessorDate.Int64,
		LessorAddress: contractModel.LessorAddress.String,
		LessorName:    contractModel.LessorName.String,
		// Room: types.Room{
		// 	RoomID:   roomModel.Id,
		// 	Name:     roomModel.Name.String,
		// 	Status:   roomModel.Status,
		// 	Capacity: roomModel.Capacity.Int64,
		// 	EIndex:   roomModel.EIndex.Int64,
		// 	WIndex:   roomModel.WIndex.Int64,
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

	l.Logger.Info("CreateContractLogic Success: ", userID)
	return &types.CreateContractRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Contract: contract,
	}, nil
}
