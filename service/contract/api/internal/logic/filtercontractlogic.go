package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
	"roomrover/service/contract/model"
	inventoryModel "roomrover/service/inventory/model"

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
	var listRoomID []int64
	var houseIDs []int64

	var mapUserPhone = map[int64]string{}
	var mapRoom = map[int64]inventoryModel.RoomTbl{}
	var mapHouse = map[int64]inventoryModel.HouseTbl{}

	var listContract []types.Contract

	var contractModels []*model.ContractTbl
	var roomModels []*inventoryModel.RoomTbl
	var houseModels []*inventoryModel.HouseTbl

	var lessorID, renterID int64

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
	userModel, err := l.svcCtx.AccountFunction.GetUserByID(userID)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterContractRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}
	if userModel.Role.Int64 == common.USER_ROLE_LESSOR {
		lessorID = userModel.Id
	} else {
		renterID = userModel.Id
	}

	count, err = l.svcCtx.ContractModel.CountContractByCondition(l.ctx, lessorID, renterID, req.Search, req.Status, req.CreateFrom, req.CreateTo)
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

	contractModels, err = l.svcCtx.ContractModel.FindContractByCondition(l.ctx, lessorID, renterID, req.Search, req.Status, req.CreateFrom, req.CreateTo, req.Offset, req.Limit)
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
		listRoomID = append(listRoomID, contractModel.RoomId.Int64)
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

	roomModels, err = l.svcCtx.InventFunction.GetRoomsByIDs(listRoomID)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, roomModel := range roomModels {
		mapRoom[roomModel.Id] = *roomModel
		houseIDs = append(houseIDs, roomModel.HouseId.Int64)
	}

	houseModels, err = l.svcCtx.InventFunction.GetHousesByIDs(houseIDs)
	if err != nil {
		l.Logger.Error(err)
		return &types.FilterContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, houseModel := range houseModels {
		mapHouse[houseModel.Id] = *houseModel
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
			ContractID: contractModel.Id,
			Code:       contractModel.Code.String,
			Status:     contractModel.Status.Int64,
			Renter: types.User{
				UserID:   contractModel.RenterId.Int64,
				Phone:    mapUserPhone[contractModel.RenterId.Int64],
				FullName: contractModel.RenterName.String,
			},
			Lessor: types.User{
				UserID:   contractModel.LessorId.Int64,
				Phone:    mapUserPhone[contractModel.LessorId.Int64],
				FullName: contractModel.LessorName.String,
			},
			Room: types.Room{
				RoomID:    contractModel.RoomId.Int64,
				HouseID:   mapRoom[contractModel.RoomId.Int64].HouseId.Int64,
				HouseName: mapHouse[mapRoom[contractModel.RoomId.Int64].HouseId.Int64].Name.String,
				Area:      mapHouse[mapRoom[contractModel.RoomId.Int64].HouseId.Int64].Area,
				Price:     mapHouse[mapRoom[contractModel.RoomId.Int64].HouseId.Int64].Price,
				Type:      mapHouse[mapRoom[contractModel.RoomId.Int64].HouseId.Int64].Type,
				Name:      mapRoom[contractModel.RoomId.Int64].Name.String,
				Status:    mapRoom[contractModel.RoomId.Int64].Status,
				Capacity:  mapRoom[contractModel.RoomId.Int64].Capacity.Int64,
				EIndex:    mapRoom[contractModel.RoomId.Int64].EIndex.Int64,
				WIndex:    mapRoom[contractModel.RoomId.Int64].WIndex.Int64,
			},
			CheckIn:  contractModel.CheckIn.Int64,
			Duration: contractModel.Duration.Int64,
			Purpose:  contractModel.Purpose.String,
			Payment: types.Payment{
				PaymentID:   paymentModel.Id,
				ContractID:  paymentModel.ContractId,
				Amount:      paymentModel.Amount,
				Discount:    paymentModel.Discount,
				Deposit:     paymentModel.Deposit,
				DepositDate: paymentModel.DepositDate,
				NextBill:    paymentModel.NextBill,
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
