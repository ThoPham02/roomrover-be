package logic

import (
	"context"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomLogic {
	return &GetRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoomLogic) GetRoom(req *types.GetRoomReq) (resp *types.GetRoomRes, err error) {
	l.Logger.Info("GetRoom: ", req)

	var userID int64
	var albums []string
	var services []types.Service
	var lessor types.User
	var contract types.Contract
	var bills []types.Bill

	var roomModel *model.RoomTbl
	var houseModel *model.HouseTbl

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetRoomRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	roomModel, err = l.svcCtx.RoomModel.FindOne(l.ctx, req.ID)
	if err != nil {
		if err == model.ErrNotFound || roomModel == nil {
			return &types.GetRoomRes{
				Result: types.Result{
					Code:    common.ROOM_NOT_FOUND_CODE,
					Message: common.ROOM_NOT_FOUND_MESS,
				},
			}, nil
		}
		l.Logger.Error(err)
		return &types.GetRoomRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	houseModel, err = l.svcCtx.HouseModel.FindOne(l.ctx, roomModel.HouseId.Int64)
	if err != nil || houseModel == nil {
		l.Logger.Error(err)
		return &types.GetRoomRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	albumModels, err := l.svcCtx.AlbumModel.FindByHouseID(l.ctx, houseModel.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, albumModel := range albumModels {
		albums = append(albums, albumModel.Url.String)
	}

	serviceModels, err := l.svcCtx.ServiceModel.FindByHouseID(l.ctx, houseModel.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, serviceModel := range serviceModels {
		services = append(services, types.Service{
			ServiceID: serviceModel.Id,
			HouseID:   serviceModel.HouseId.Int64,
			Name:      serviceModel.Name.String,
			Price:     serviceModel.Price.Int64,
			Unit:      serviceModel.Unit.Int64,
		})
	}

	contractModel, err := l.svcCtx.ContractFunction.GetActiveContractByRoomID(roomModel.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return &types.GetRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	if contractModel != nil {
		paymentModel, err := l.svcCtx.ContractFunction.GetPaymentByContractID(contractModel.Id)
		if err != nil || paymentModel == nil {
			l.Logger.Error(err)
			return &types.GetRoomRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		renter, err := l.svcCtx.AccountFunction.GetUserByID(contractModel.RenterId.Int64)
		if err != nil || renter == nil {
			l.Logger.Error(err)
			return &types.GetRoomRes{
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
			Lessor: types.User{
				UserID:      contractModel.LessorId.Int64,
				FullName:    contractModel.LessorName.String,
				CccdNumber:  contractModel.LessorNumber.String,
				CccdDate:    contractModel.LessorDate.Int64,
				CccdAddress: contractModel.LessorAddress.String,
			},
			Renter: types.User{
				UserID:      renter.Id,
				Phone:       renter.Phone,
				FullName:    renter.FullName.String,
				CccdNumber:  renter.CCCDNumber.String,
				CccdDate:    renter.CCCDDate.Int64,
				CccdAddress: renter.CCCDAddress.String,
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
			CreatedBy: contractModel.CreatedBy.Int64,
			UpdatedBy: contractModel.UpdatedBy.Int64,
		}
	}

	contractModels, err := l.svcCtx.ContractFunction.GetContractByRoom(roomModel.Id)
	if err != nil {
		l.Logger.Error(err)
		return &types.GetRoomRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	for _, contract := range contractModels {
		renter, err := l.svcCtx.AccountFunction.GetUserByID(contract.RenterId.Int64)
		if err != nil {
			l.Logger.Error(err)
			return &types.GetRoomRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		billModels, err := l.svcCtx.ContractFunction.GetBillByContractID(contract.Id)
		if err != nil {
			l.Logger.Error(err)
			return &types.GetRoomRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}

		for _, bill := range billModels {
			bills = append(bills, types.Bill{
				BillID:       bill.Id,
				Title:        bill.Title.String,
				ContractCode: contract.Code.String,
				RenterID:     renter.Id,
				RenterName:   renter.FullName.String,
				RenterPhone:  renter.Phone,
				PaymentID:    bill.PaymentId,
				PaymentDate:  bill.PaymentDate.Int64,
				Amount:       bill.Amount,
				Remain:       bill.Remain,
				Status:       bill.Status,
			})
		}
	}

	l.Logger.Info("GetRoom Success: ", userID)
	return &types.GetRoomRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		Room: types.Room{
			RoomID:    roomModel.Id,
			HouseID:   roomModel.HouseId.Int64,
			Name:      roomModel.Name.String,
			HouseName: houseModel.Name.String,
			Area:      houseModel.Area,
			Price:     houseModel.Price,
			Type:      houseModel.Type,
			Status:    roomModel.Status,
			Capacity:  roomModel.Capacity.Int64,
			EIndex:    roomModel.EIndex.Int64,
			WIndex:    roomModel.WIndex.Int64,
		},
		House: types.House{
			HouseID:     houseModel.Id,
			User:        lessor,
			Name:        houseModel.Name.String,
			Description: houseModel.Description.String,
			Type:        houseModel.Type,
			Status:      houseModel.Status,
			Area:        houseModel.Area,
			Price:       houseModel.Price,
			BedNum:      houseModel.BedNum.Int64,
			LivingNum:   houseModel.LivingNum.Int64,
			Albums:      albums,
			Services:    services,
			Address:     houseModel.Address.String,
			WardID:      houseModel.WardId,
			DistrictID:  houseModel.DistrictId,
			ProvinceID:  houseModel.ProvinceId,
			CreatedAt:   houseModel.CreatedAt.Int64,
			UpdatedAt:   houseModel.UpdatedAt.Int64,
			CreatedBy:   houseModel.CreatedBy.Int64,
			UpdatedBy:   houseModel.UpdatedBy.Int64,
		},
		Contract: contract,
		Bills:    bills,
	}, nil
}
