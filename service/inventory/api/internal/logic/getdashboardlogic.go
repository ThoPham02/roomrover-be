package logic

import (
	"context"
	"fmt"

	"roomrover/common"
	"roomrover/service/inventory/api/internal/svc"
	"roomrover/service/inventory/api/internal/types"
	"roomrover/service/inventory/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDashboardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDashboardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDashboardLogic {
	return &GetDashboardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDashboardLogic) GetDashboard(req *types.GetDashboardReq) (resp *types.GetDashboardRes, err error) {
	l.Logger.Info("GetDashboard: ", req)

	var userID int64

	var houseModels []*model.HouseTbl
	var mapHouseName = make(map[int64]string)

	var totalRoom int
	var rentedRoom int
	var totalAmount int64

	var contacts []types.Contact
	var contracts []types.Contract
	var houseRevenue []types.HouseRevenue
	var mapHouseAmount = make(map[int64]int64)

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Info(err)
		return &types.GetDashboardRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	_, houseModels, err = l.svcCtx.HouseModel.FilterHouse(l.ctx, userID, "", 0, 0, 0, 0)
	if err != nil {
		l.Logger.Info(err)
		return &types.GetDashboardRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	for _, house := range houseModels {
		mapHouseAmount[house.Id] = 0
		mapHouseName[house.Id] = house.Name.String

		roomModels, _, err := l.svcCtx.RoomModel.FindByHouseID(l.ctx, house.Id, 0, 0)
		if err != nil {
			l.Logger.Info(err)
			return &types.GetDashboardRes{
				Result: types.Result{
					Code:    common.UNKNOWN_ERR_CODE,
					Message: common.UNKNOWN_ERR_MESS,
				},
			}, nil
		}

		for _, room := range roomModels {
			totalRoom++
			if room.Status == common.ROOM_STATUS_RENTED {
				rentedRoom++
			}

			contractModels, err := l.svcCtx.ContractFunction.GetContractByRoom(room.Id)
			if err != nil {
				l.Logger.Info(err)
				return &types.GetDashboardRes{
					Result: types.Result{
						Code:    common.UNKNOWN_ERR_CODE,
						Message: common.UNKNOWN_ERR_MESS,
					},
				}, nil
			}

			for _, contract := range contractModels {
				if contract.Status.Int64 == common.CONTRACT_STATUS_NEARLY_OUT_DATE {
					renter, err := l.svcCtx.AccountFunction.GetUserByID(contract.RenterId.Int64)
					if err != nil {
						l.Logger.Info(err)
						return &types.GetDashboardRes{
							Result: types.Result{
								Code:    common.UNKNOWN_ERR_CODE,
								Message: common.UNKNOWN_ERR_MESS,
							},
						}, nil
					}

					contracts = append(contracts, types.Contract{
						ContractID: contract.Id,
						Code:       contract.Code.String,
						Renter: types.User{
							UserID:   renter.Id,
							Phone:    renter.Phone,
							FullName: renter.FullName.String,
						},
						Room: types.Room{
							RoomID:    room.Id,
							HouseName: fmt.Sprintf("%s (%s)", room.Name.String, house.Name.String),
						},
						CheckIn: common.GetBillTimeByIndex(contract.CheckIn.Int64, int(contract.Duration.Int64)),
					})
				}

				billPayModels, err := l.svcCtx.ContractFunction.GetBillPayByContractID(contract.Id)
				if err != nil {
					l.Logger.Info(err)
					return &types.GetDashboardRes{
						Result: types.Result{
							Code:    common.UNKNOWN_ERR_CODE,
							Message: common.UNKNOWN_ERR_MESS,
						},
					}, nil
				}

				for _, billPay := range billPayModels {
					if billPay.Status == common.PAYMENT_DETAIL_STATUS_DONE {
						totalAmount += billPay.Amount
						mapHouseAmount[house.Id] += billPay.Amount
					}
				}
			}
		}
	}

	for _, house := range houseModels {
		houseRevenue = append(houseRevenue, types.HouseRevenue{
			HouseID:   house.Id,
			HouseName: house.Name.String,
			Revenue:   mapHouseAmount[house.Id],
		})
	}

	contactModels, err := l.svcCtx.ContactModel.GetCurrentContact(l.ctx, userID)
	if err != nil {
		l.Logger.Info(err)
		return &types.GetDashboardRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}
	for _, contact := range contactModels {
		renter, err := l.svcCtx.AccountFunction.GetUserByID(contact.RenterId.Int64)
		if err != nil {
			l.Logger.Info(err)
			return &types.GetDashboardRes{
				Result: types.Result{
					Code:    common.UNKNOWN_ERR_CODE,
					Message: common.UNKNOWN_ERR_MESS,
				},
			}, nil
		}

		contacts = append(contacts, types.Contact{
			ID:          contact.Id,
			HouseID:     contact.HouseId.Int64,
			HouseName:   mapHouseName[contact.HouseId.Int64],
			RenterID:    renter.Id,
			RenterName:  renter.FullName.String,
			RenterPhone: renter.Phone,
			Datetime:    contact.Datetime.Int64,
			Status:      contact.Status,
		})
	}

	l.Logger.Info("GetDashboard Success: ", userID)
	return &types.GetDashboardRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		TotalRoom:       totalRoom,
		RentedRoom:      rentedRoom,
		EmptyRoom:       totalRoom - rentedRoom,
		TotalAmount:     totalAmount,
		CurrentContact:  contacts,
		ExpiredContract: contracts,
		HouseRevenue:    houseRevenue,
	}, nil
}
