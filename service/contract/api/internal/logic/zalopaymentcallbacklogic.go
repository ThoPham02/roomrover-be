package logic

import (
	"context"
	"database/sql"
	"encoding/json"

	"roomrover/common"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ZaloPaymentCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewZaloPaymentCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ZaloPaymentCallbackLogic {
	return &ZaloPaymentCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type CallbackData struct {
	AppID          int    `json:"app_id"`
	AppTransID     string `json:"app_trans_id"`
	AppTime        int64  `json:"app_time"`
	AppUser        string `json:"app_user"`
	Amount         int    `json:"amount"`
	EmbedData      string `json:"embed_data"`
	Item           string `json:"item"`
	ZpTransID      int64  `json:"zp_trans_id"`
	ServerTime     int64  `json:"server_time"`
	Channel        int    `json:"channel"`
	MerchantUserID string `json:"merchant_user_id"`
	ZpUserID       string `json:"zp_user_id"`
	UserFeeAmount  int    `json:"user_fee_amount"`
	DiscountAmount int    `json:"discount_amount"`
}

func (l *ZaloPaymentCallbackLogic) ZaloPaymentCallback(req *types.ZaloPaymentCallbackReq) (resp *types.ZaloPaymentCallbackRes, err error) {
	l.Logger.Info("ZaloPaymentCallback: ", req)

	var data CallbackData

	err = json.Unmarshal([]byte(req.Data), &data)
	if err != nil {
		l.Logger.Error(err)
		return &types.ZaloPaymentCallbackRes{
			ReturnCode:    2,
			ReturnMessage: "Invalid data",
		}, nil
	}

	billPayModel, err := l.svcCtx.BillPayModel.FindOneByTransID(l.ctx, data.AppTransID)
	if err != nil || billPayModel == nil {
		l.Logger.Error(err)
		return &types.ZaloPaymentCallbackRes{
			ReturnCode:    2,
			ReturnMessage: "Invalid data",
		}, nil
	}

	billModel, err := l.svcCtx.BillModel.FindOne(l.ctx, billPayModel.BillId)
	if err != nil || billModel == nil {
		l.Logger.Error(err)
		return &types.ZaloPaymentCallbackRes{
			ReturnCode:    2,
			ReturnMessage: "Invalid data",
		}, nil
	}

	billModel.Remain = billModel.Remain - int64(data.Amount)
	if billModel.Remain <= 0 {
		billModel.Status = common.BILL_STATUS_PAID
		billModel.PaidDate = sql.NullInt64{Valid: true, Int64: common.GetCurrentTime()}
	}
	err = l.svcCtx.BillModel.Update(l.ctx, billModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.ZaloPaymentCallbackRes{
			ReturnCode:    2,
			ReturnMessage: "Invalid data",
		}, nil
	}
	billPayModel.Status = common.BILL_PAY_STATUS_DONE
	err = l.svcCtx.BillPayModel.Update(l.ctx, billPayModel)
	if err != nil {
		l.Logger.Error(err)
		return &types.ZaloPaymentCallbackRes{
			ReturnCode:    2,
			ReturnMessage: "Invalid data",
		}, nil
	}

	l.Logger.Info("ZaloPaymentCallback Success: ", data.AppTransID)
	return &types.ZaloPaymentCallbackRes{
		ReturnCode:    1,
		ReturnMessage: "Success",
	}, nil
}
