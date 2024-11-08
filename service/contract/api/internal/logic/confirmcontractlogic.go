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

type ConfirmContractLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfirmContractLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmContractLogic {
	return &ConfirmContractLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfirmContractLogic) ConfirmContract(req *types.ConfirmContractReq) (resp *types.ConfirmContractRes, err error) {
	l.Logger.Info("ConfirmContract", req)

	var userID int64
	var albums []string

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.ConfirmContractRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	contractModel, err := l.svcCtx.ContractModel.FindOne(l.ctx, req.ID)
	if err != nil || contractModel == nil {
		l.Logger.Error(err)
		if err == model.ErrNotFound {
			return &types.ConfirmContractRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		return &types.ConfirmContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	paymentModel, err := l.svcCtx.PaymentModel.FindByContractID(l.ctx, contractModel.Id)
	if err != nil || paymentModel == nil {
		l.Logger.Error(err)
		if err == model.ErrNotFound {
			return &types.ConfirmContractRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
		return &types.ConfirmContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}
	var paymentID = paymentModel.Id

	if req.Albums != "" {
		err = json.Unmarshal([]byte(req.Albums), &albums)
		if err != nil {
			l.Logger.Error(err)
			return &types.ConfirmContractRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}
	}
	if req.Renters != "" {
		var renters []types.PaymentRenter
		var updateRenters []*model.PaymentRenterTbl
		err = json.Unmarshal([]byte(req.Renters), &renters)
		if err != nil {
			l.Logger.Error(err)
			return &types.ConfirmContractRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}

		for _, renter := range renters {
			var userID int64
			userModel, err := l.svcCtx.AccountFunction.FindUserByPhone(renter.Phone)
			if err != nil {
				l.Logger.Error(err)
				return &types.ConfirmContractRes{
					Result: types.Result{
						Code:    common.DB_ERR_CODE,
						Message: common.DB_ERR_MESS,
					},
				}, nil
			}
			if userModel == nil {
				userID = l.svcCtx.ObjSync.GenServiceObjID()
				err = l.svcCtx.AccountFunction.CreateInactivatedUser(userID, renter.Phone, renter.Name, renter.CccdNumber, renter.CccdDate, renter.CccdAddress)
				if err != nil {
					l.Logger.Error(err)
					return &types.ConfirmContractRes{
						Result: types.Result{
							Code:    common.DB_ERR_CODE,
							Message: common.DB_ERR_MESS,
						},
					}, nil
				}
			} else {
				userID = userModel.Id
				userModel.FullName = sql.NullString{Valid: true, String: renter.Name}
				userModel.CCCDNumber = sql.NullString{Valid: true, String: renter.CccdNumber}
				userModel.CCCDDate = sql.NullInt64{Valid: true, Int64: renter.CccdDate}
				userModel.CCCDAddress = sql.NullString{Valid: true, String: renter.CccdAddress}

				err = l.svcCtx.AccountFunction.UpdateUser(userModel)
				if err != nil {
					l.Logger.Error(err)
					return &types.ConfirmContractRes{
						Result: types.Result{
							Code:    common.DB_ERR_CODE,
							Message: common.DB_ERR_MESS,
						},
					}, nil
				}
			}

			if renter.ID < common.MIN_ID {
				updateRenters = append(updateRenters, &model.PaymentRenterTbl{
					Id:        l.svcCtx.ObjSync.GenServiceObjID(),
					PaymentId: sql.NullInt64{Valid: true, Int64: paymentID},
					UserId:    sql.NullInt64{Valid: true, Int64: userID},
				})
			} else {
				renterModel, err := l.svcCtx.PaymentRenterModel.FindOne(l.ctx, renter.ID)
				if err != nil || renterModel == nil {
					l.Logger.Error(err)
					return &types.ConfirmContractRes{
						Result: types.Result{
							Code:    common.DB_ERR_CODE,
							Message: common.DB_ERR_MESS,
						},
					}, nil
				}

				renterModel.UserId = sql.NullInt64{Valid: true, Int64: userID}
				err = l.svcCtx.PaymentRenterModel.Update(l.ctx, renterModel)
				if err != nil {
					l.Logger.Error(err)
					return &types.ConfirmContractRes{
						Result: types.Result{
							Code:    common.DB_ERR_CODE,
							Message: common.DB_ERR_MESS,
						},
					}, nil
				}

				updateRenters = append(updateRenters, renterModel)
				continue
			}
		}

		err = l.svcCtx.PaymentRenterModel.DeleteByPaymentID(l.ctx, paymentID)
		if err != nil {
			l.Logger.Error(err)
			return &types.ConfirmContractRes{
				Result: types.Result{
					Code:    common.DB_ERR_CODE,
					Message: common.DB_ERR_MESS,
				},
			}, nil
		}
		for _, renter := range updateRenters {
			_, err = l.svcCtx.PaymentRenterModel.Insert(l.ctx, renter)
			if err != nil {
				l.Logger.Error(err)
				return &types.ConfirmContractRes{
					Result: types.Result{
						Code:    common.DB_ERR_CODE,
						Message: common.DB_ERR_MESS,
					},
				}, nil
			}
		}
	}
	if req.Services != "" {
		var services []types.PaymentDetail
		err = json.Unmarshal([]byte(req.Services), &services)
		if err != nil {
			l.Logger.Error(err)
			return &types.ConfirmContractRes{
				Result: types.Result{
					Code:    common.INVALID_REQUEST_CODE,
					Message: common.INVALID_REQUEST_MESS,
				},
			}, nil
		}

		for _, service := range services {
			detailModel, err := l.svcCtx.PaymentDetailModel.FindOne(l.ctx, service.ID)
			if err != nil || detailModel == nil {
				l.Logger.Error(err)
				return &types.ConfirmContractRes{
					Result: types.Result{
						Code:    common.DB_ERR_CODE,
						Message: common.DB_ERR_MESS,
					},
				}, nil
			}

			detailModel.PaymentId = sql.NullInt64{Valid: true, Int64: paymentID}
			detailModel.Index = sql.NullInt64{Valid: true, Int64: service.Index}
			err = l.svcCtx.PaymentDetailModel.Update(l.ctx, detailModel)
			if err != nil {
				l.Logger.Error(err)
				return &types.ConfirmContractRes{
					Result: types.Result{
						Code:    common.DB_ERR_CODE,
						Message: common.DB_ERR_MESS,
					},
				}, nil
			}
		}
	}

	renterModel, err := l.svcCtx.AccountFunction.GetUserByID(contractModel.RenterId.Int64)
	if err != nil || renterModel == nil {
		l.Logger.Error(err)
		return &types.ConfirmContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	contractModel.Status = sql.NullInt64{Valid: true, Int64: common.CONTRACT_STATUS_ACTIVE}
	contractModel.ConfirmedImgs = sql.NullString{Valid: true, String: req.Albums}
	contractModel.UpdatedAt = sql.NullInt64{Valid: true, Int64: common.GetCurrentTime()}
	contractModel.UpdatedBy = sql.NullInt64{Valid: true, Int64: userID}
	contractModel.RenterId = sql.NullInt64{Valid: true, Int64: renterModel.Id}
	contractModel.RenterName = sql.NullString{Valid: true, String: renterModel.FullName.String}
	contractModel.RenterNumber = sql.NullString{Valid: true, String: renterModel.CCCDNumber.String}
	contractModel.RenterDate = sql.NullInt64{Valid: true, Int64: renterModel.CCCDDate.Int64}
	contractModel.RenterAddress = sql.NullString{Valid: true, String: renterModel.CCCDAddress.String}
	err = l.svcCtx.ContractModel.Update(l.ctx, contractModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.ConfirmContractRes{
			Result: types.Result{
				Code:    common.DB_ERR_CODE,
				Message: common.DB_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("ConfirmContract Success", userID)
	return &types.ConfirmContractRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
	}, nil
}
