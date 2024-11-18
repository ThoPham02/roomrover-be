package logic

import (
	"context"
	"database/sql"
	"encoding/json"

	"roomrover/common"
	accountModel "roomrover/service/account/model"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
	model "roomrover/service/contract/model"
	inventoryModel "roomrover/service/inventory/model"
	notiModel "roomrover/service/notification/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UpdateContractLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateContractLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContractLogic {
	return &UpdateContractLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateContractLogic) UpdateContract(req *types.UpdateContractReq) (resp *types.UpdateContractRes, err error) {
	l.Logger.Info("UpdateContract", req)

	var userID int64
	var currentTime = common.GetCurrentTime()
	var contract types.Contract
	var paymentRenters []types.PaymentRenter
	var paymentDetails []types.PaymentDetail
	var renter types.User
	var lessor types.User
	var room types.Room

	var renterModel *accountModel.UserTbl
	var lessorModel *accountModel.UserTbl
	var roomModel *inventoryModel.RoomTbl

	var serviceModels []*inventoryModel.ServiceTbl
	var contractModel *model.ContractTbl
	var paymentModel *model.PaymentTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	contractModel, err = l.svcCtx.ContractModel.FindOne(l.ctx, req.ID)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return &types.UpdateContractRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	paymentModel, err = l.svcCtx.PaymentModel.FindByContractID(l.ctx, contractModel.Id)
	if err != nil || paymentModel == nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	err = json.Unmarshal([]byte(req.Renter), &renter)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}
	err = json.Unmarshal([]byte(req.Lessor), &lessor)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}
	err = json.Unmarshal([]byte(req.Room), &room)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}

	renterModel, err = l.svcCtx.AccountFunction.GetUserByID(renter.UserID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	lessorModel, err = l.svcCtx.AccountFunction.GetUserByID(lessor.UserID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	roomModel, err = l.svcCtx.InventFunction.GetRoomByID(room.RoomID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if renterModel == nil || lessorModel == nil || roomModel == nil || (roomModel.Status != common.ROOM_STATUS_ACTIVE && room.RoomID != contractModel.RoomId.Int64) {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.INVALID_REQUEST_CODE,
				Message: common.INVALID_REQUEST_MESS,
			},
		}, nil
	}
	if room.RoomID != contractModel.RoomId.Int64 {
		oldRoom, err := l.svcCtx.InventFunction.GetRoomByID(contractModel.RoomId.Int64)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		oldRoom.Status = common.ROOM_STATUS_ACTIVE
		err = l.svcCtx.InventFunction.UpdateRoom(oldRoom)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		houseModel, err := l.svcCtx.InventFunction.GetHouseByID(oldRoom.HouseId.Int64)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
		houseModel.Status = common.HOUSE_STATUS_ACTIVE
		err = l.svcCtx.InventFunction.UpdateHouse(houseModel)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}

	serviceModels, err = l.svcCtx.InventFunction.GetSericesByRoom(room.RoomID)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	renterModel.CCCDNumber = sql.NullString{String: renter.CccdNumber, Valid: true}
	renterModel.CCCDDate = sql.NullInt64{Int64: renter.CccdDate, Valid: true}
	renterModel.CCCDAddress = sql.NullString{String: renter.CccdAddress, Valid: true}
	renterModel.FullName = sql.NullString{String: renter.FullName, Valid: true}
	renterModel.UpdatedAt = sql.NullInt64{Int64: currentTime, Valid: true}
	err = l.svcCtx.AccountFunction.UpdateUser(renterModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	lessorModel.CCCDNumber = sql.NullString{String: lessor.CccdNumber, Valid: true}
	lessorModel.CCCDDate = sql.NullInt64{Int64: lessor.CccdDate, Valid: true}
	lessorModel.CCCDAddress = sql.NullString{String: lessor.CccdAddress, Valid: true}
	lessorModel.FullName = sql.NullString{String: lessor.FullName, Valid: true}
	lessorModel.UpdatedAt = sql.NullInt64{Int64: currentTime, Valid: true}
	err = l.svcCtx.AccountFunction.UpdateUser(lessorModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	roomActive, err := l.svcCtx.InventFunction.CountRoomActiveByHouseID(roomModel.HouseId.Int64)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	if roomActive == 1 {
		houseModel, err := l.svcCtx.InventFunction.GetHouseByID(roomModel.HouseId.Int64)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
		houseModel.Status = common.HOUSE_STATUS_SOLD_OUT
		houseModel.UpdatedAt = sql.NullInt64{Int64: currentTime, Valid: true}
		err = l.svcCtx.InventFunction.UpdateHouse(houseModel)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}

	contractModel.RenterId = sql.NullInt64{Valid: true, Int64: renterModel.Id}
	contractModel.RenterNumber = sql.NullString{Valid: true, String: renter.CccdNumber}
	contractModel.RenterDate = sql.NullInt64{Valid: true, Int64: renter.CccdDate}
	contractModel.RenterAddress = sql.NullString{Valid: true, String: renter.CccdAddress}
	contractModel.RenterName = sql.NullString{Valid: true, String: renter.FullName}
	contractModel.LessorId = sql.NullInt64{Valid: true, Int64: lessorModel.Id}
	contractModel.LessorNumber = sql.NullString{Valid: true, String: lessor.CccdNumber}
	contractModel.LessorDate = sql.NullInt64{Valid: true, Int64: lessor.CccdDate}
	contractModel.LessorAddress = sql.NullString{Valid: true, String: lessor.CccdAddress}
	contractModel.LessorName = sql.NullString{Valid: true, String: lessor.FullName}
	contractModel.RoomId = sql.NullInt64{Valid: true, Int64: roomModel.Id}
	contractModel.CheckIn = sql.NullInt64{Valid: true, Int64: req.CheckIn}
	contractModel.Duration = sql.NullInt64{Valid: true, Int64: req.Duration}
	contractModel.Purpose = sql.NullString{Valid: true, String: req.Purpose}
	contractModel.UpdatedAt = sql.NullInt64{Valid: true, Int64: currentTime}
	contractModel.UpdatedBy = sql.NullInt64{Valid: true, Int64: userID}

	paymentModel.Amount = req.Price
	paymentModel.Discount = req.Discount
	paymentModel.Deposit = req.Deposit
	paymentModel.DepositDate = req.DepositDate
	paymentModel.NextBill = common.GetNextMonthDate(req.CheckIn)

	roomModel.EIndex = sql.NullInt64{Int64: room.EIndex, Valid: true}
	roomModel.WIndex = sql.NullInt64{Int64: room.WIndex, Valid: true}
	roomModel.Status = common.ROOM_STATUS_RENTED
	err = l.svcCtx.InventFunction.UpdateRoom(roomModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.PaymentDetailModel.DeleteByPaymentID(l.ctx, paymentModel.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, service := range serviceModels {
		paymentDetails = append(paymentDetails, types.PaymentDetail{
			ID:        userID,
			PaymentID: paymentModel.Id,
			Name:      service.Name.String,
			Price:     service.Price.Int64,
			Type:      service.Unit.Int64,
		})

		paymentDetail := &model.PaymentDetailTbl{
			Id:        l.svcCtx.ObjSync.GenServiceObjID(),
			PaymentId: sql.NullInt64{Valid: true, Int64: paymentModel.Id},
			Name:      sql.NullString{Valid: true, String: service.Name.String},
			Type:      sql.NullInt64{Valid: true, Int64: service.Unit.Int64},
			Price:     sql.NullInt64{Valid: true, Int64: service.Price.Int64},
		}
		_, err = l.svcCtx.PaymentDetailModel.Insert(l.ctx, paymentDetail)
		if err != nil {
			l.Logger.Error(err)
			return &types.UpdateContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
	}

	err = l.svcCtx.ContractModel.Update(l.ctx, contractModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	err = l.svcCtx.PaymentModel.Update(l.ctx, paymentModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	err = l.svcCtx.NotiFunction.CreateNotification(&notiModel.NotificationTbl{
		Id:        l.svcCtx.ObjSync.GenServiceObjID(),
		Sender:    lessorModel.Id,
		Receiver:  renterModel.Id,
		RefId:     contractModel.Id,
		RefType:   common.NOTI_TYPE_UPDATE_CONTRACT,
		Unread:    common.NOTI_TYPE_UNREAD,
		CreatedAt: currentTime,
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.UpdateContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	contract = types.Contract{
		ContractID: contractModel.Id,
		Code:       contractModel.Code.String,
		Status:     contractModel.Status.Int64,
		Room:       room,
		CheckIn:    contractModel.CheckIn.Int64,
		Duration:   contractModel.Duration.Int64,
		Purpose:    contractModel.Purpose.String,
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

	l.Logger.Info("CreateContractLogic Success: ", userID)
	return &types.UpdateContractRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Contract: contract,
	}, nil
}
