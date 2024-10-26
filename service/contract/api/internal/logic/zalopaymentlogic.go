package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"roomrover/common"
	"roomrover/service/contract/api/internal/svc"
	"roomrover/service/contract/api/internal/types"
	"roomrover/service/contract/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zpmep/hmacutil"
)

type ZaloPaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewZaloPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ZaloPaymentLogic {
	return &ZaloPaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type object map[string]interface{}

type ZaloRes struct {
	ReturnCode       int    `json:"return_code"`
	ReturnMessage    string `json:"return_message"`
	SubReturnCode    int    `json:"sub_return_code"`
	SubReturnMessage string `json:"sub_return_message"`
	ZpTransToken     string `json:"zp_trans_token"`
	OrderUrl         string `json:"order_url"`
	QrCode           string `json:"qr_code"`
}

func (l *ZaloPaymentLogic) ZaloPayment(req *types.ZaloPaymentReq) (resp *types.ZaloPaymentRes, err error) {
	l.Logger.Info("ZaloPayment: ", req)

	var userID int64
	var redirecPath string

	userID, err = common.GetUserIDFromContext(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return &types.ZaloPaymentRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	userModel, err := l.svcCtx.AccountFunction.GetUserByID(userID)
	if err != nil || userModel == nil {
		l.Logger.Error(err)
		return &types.ZaloPaymentRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	if userModel.Role.Int64 == common.USER_ROLE_RENTER {
		redirecPath = "/renter/payment"
	} else {
		redirecPath = "/lessor/payment"
	}

	billModel, err := l.svcCtx.BillModel.FindOne(l.ctx, req.BillID)
	if err != nil || billModel == nil {
		l.Logger.Error(err)
		return &types.ZaloPaymentRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}
	paymentModel, err := l.svcCtx.PaymentModel.FindOne(l.ctx, billModel.PaymentId)
	if err != nil || paymentModel == nil {
		l.Logger.Error(err)
		return &types.ZaloPaymentRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}
	contractModel, err := l.svcCtx.ContractModel.FindOne(l.ctx, paymentModel.ContractId)
	if err != nil || contractModel == nil {
		l.Logger.Error(err)
		return &types.ZaloPaymentRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}
	lessorModel, err := l.svcCtx.AccountFunction.GetUserByID(contractModel.LessorId.Int64)
	if err != nil || lessorModel == nil {
		l.Logger.Error(err)
		return &types.ZaloPaymentRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	// generate Zalo payment code and QR code
	billPayID := l.svcCtx.ObjSync.GenServiceObjID()
	transID := strconv.FormatInt(billPayID, 10)
	embedData, _ := json.Marshal(object{
		"redirecturl": l.svcCtx.Config.ZaloPay.RedirectDomain + redirecPath,
	})
	items, _ := json.Marshal([]object{})

	params := make(url.Values)
	params.Add("app_id", l.svcCtx.Config.ZaloPay.AppID)
	params.Add("amount", strconv.FormatInt(int64(billModel.Amount), 10))
	params.Add("app_user", "user123")
	params.Add("embed_data", string(embedData))
	params.Add("item", string(items))
	params.Add("description", "RoomRover - Thanh toan hoa don thue nha ")
	params.Add("bank_code", l.svcCtx.Config.ZaloPay.BankCode)
	params.Add("phone", lessorModel.Phone)
	params.Add("email", "thoahlgbg2002@gmail.com")
	params.Add("address", lessorModel.Address.String)
	params.Add("callback_url", l.svcCtx.Config.ZaloPay.CallbackUrl)

	now := time.Now()
	params.Add("app_time", strconv.FormatInt(now.UnixNano()/int64(time.Millisecond), 10))
	params.Add("app_trans_id", fmt.Sprintf("%02d%02d%02d_%v", now.Year()%100, int(now.Month()), now.Day(), transID))

	data := fmt.Sprintf("%v|%v|%v|%v|%v|%v|%v", params.Get("app_id"), params.Get("app_trans_id"), params.Get("app_user"),
		params.Get("amount"), params.Get("app_time"), params.Get("embed_data"), params.Get("item"))
	params.Add("mac", hmacutil.HexStringEncode(hmacutil.SHA256, l.svcCtx.Config.ZaloPay.Key1, data))

	res, err := http.PostForm("https://sb-openapi.zalopay.vn/v2/create", params)
	if err != nil {
		return &types.ZaloPaymentRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var result ZaloRes

	if err := json.Unmarshal(body, &result); err != nil {
		return &types.ZaloPaymentRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}
	if result.ReturnCode != 1 {
		return &types.ZaloPaymentRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	_, err = l.svcCtx.BillPayModel.Insert(l.ctx, &model.BillPayTbl{
		Id:      billPayID,
		BillId:  billModel.Id,
		UserId:  userID,
		Amount:  billModel.Amount,
		PayDate: common.GetCurrentTime(),
		Status:  common.BILL_PAY_STATUS_PROCESS,
		TransId: sql.NullString{Valid: true, String: time.Now().Format("060102") + "_" + transID},
		Type:    common.BILL_PAY_TYPE_ZALO,
		Url:     sql.NullString{},
	})
	if err != nil {
		l.Logger.Error(err)
		return &types.ZaloPaymentRes{
			Result: types.Result{
				Code:    common.UNKNOWN_ERR_CODE,
				Message: common.UNKNOWN_ERR_MESS,
			},
		}, nil
	}

	l.Logger.Info("ZaloPayment Success: ", userID)
	return &types.ZaloPaymentRes{
		Result: types.Result{
			Code:    common.SUCCESS_CODE,
			Message: common.SUCCESS_MESS,
		},
		OrderUrl: result.OrderUrl,
		QrCode:   result.QrCode,
	}, nil
}
