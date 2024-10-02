package logic

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"roomrover/service/payment/api/internal/svc"
	"roomrover/service/payment/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MomoPaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMomoPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MomoPaymentLogic {
	return &MomoPaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Payload struct {
	PartnerCode  string `json:"partnerCode"`
	AccessKey    string `json:"accessKey"`
	RequestID    string `json:"requestId"`
	Amount       string `json:"amount"`
	OrderID      string `json:"orderId"`
	OrderInfo    string `json:"orderInfo"`
	PartnerName  string `json:"partnerName"`
	StoreId      string `json:"storeId"`
	OrderGroupId string `json:"orderGroupId"`
	Lang         string `json:"lang"`
	AutoCapture  bool   `json:"autoCapture"`
	RedirectUrl  string `json:"redirectUrl"`
	IpnUrl       string `json:"ipnUrl"`
	ExtraData    string `json:"extraData"`
	RequestType  string `json:"requestType"`
	Signature    string `json:"signature"`
}

func (l *MomoPaymentLogic) MomoPayment(req *types.MomoPaymentReq) (resp *types.MomoPaymentRes, err error) {
	l.Logger.Info("MomoPaymentLogic Start: ", req)

	var endpoint = l.svcCtx.Config.Momo.Endpoint
	var accessKey = l.svcCtx.Config.Momo.AccessKey
	var secretKey = l.svcCtx.Config.Momo.SecretKey
	var orderInfo = l.svcCtx.Config.Momo.OrderInfo
	var partnerCode = l.svcCtx.Config.Momo.PartnerCode
	var redirectUrl = l.svcCtx.Config.Momo.RedirectUrl
	var ipnUrl = l.svcCtx.Config.Momo.IpnUrl
	var amount = req.Amount
	var orderId = req.BillID
	var requestId = req.BillID
	var extraData = ""
	var partnerName = l.svcCtx.Config.Momo.PartnerName
	var storeId = l.svcCtx.Config.Momo.StoreId
	var orderGroupId = ""
	var autoCapture = l.svcCtx.Config.Momo.AutoCapture
	var lang = l.svcCtx.Config.Momo.Lang
	var requestType = l.svcCtx.Config.Momo.RequestType

	// rawSignature = "accessKey=" + accessKey + "&amount=" + amount + "&extraData=" + extraData + "&ipnUrl=" + ipnUrl \
	//				+ "&orderId=" + orderId + "&orderInfo=" + orderInfo + "&partnerCode=" + partnerCode \
	//            	+ "&redirectUrl=" + redirectUrl + "&requestId=" + requestId + "&requestType=" + requestType\

	var rawSignature bytes.Buffer
	rawSignature.WriteString("accessKey=")
	rawSignature.WriteString(accessKey)
	rawSignature.WriteString("&amount=")
	rawSignature.WriteString(fmt.Sprintf("%d", amount))
	rawSignature.WriteString("&extraData=")
	rawSignature.WriteString(extraData)
	rawSignature.WriteString("&ipnUrl=")
	rawSignature.WriteString(ipnUrl)
	rawSignature.WriteString("&orderId=")
	rawSignature.WriteString(orderId)
	rawSignature.WriteString("&orderInfo=")
	rawSignature.WriteString(orderInfo)
	rawSignature.WriteString("&partnerCode=")
	rawSignature.WriteString(partnerCode)
	rawSignature.WriteString("&redirectUrl=")
	rawSignature.WriteString(redirectUrl)
	rawSignature.WriteString("&requestId=")
	rawSignature.WriteString(requestId)
	rawSignature.WriteString("&requestType=")
	rawSignature.WriteString(requestType)

	hmac := hmac.New(sha256.New, []byte(secretKey))

	hmac.Write(rawSignature.Bytes())
	fmt.Println("Raw signature: " + rawSignature.String())

	signature := hex.EncodeToString(hmac.Sum(nil))

	var payload = Payload{
		PartnerCode:  partnerCode,
		AccessKey:    accessKey,
		RequestID:    requestId,
		Amount:       fmt.Sprintf("%d", amount),
		RequestType:  requestType,
		RedirectUrl:  redirectUrl,
		IpnUrl:       ipnUrl,
		OrderID:      orderId,
		StoreId:      storeId,
		PartnerName:  partnerName,
		OrderGroupId: orderGroupId,
		AutoCapture:  autoCapture,
		Lang:         lang,
		OrderInfo:    orderInfo,
		ExtraData:    extraData,
		Signature:    signature,
	}

	var jsonPayload []byte
	jsonPayload, err = json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Payload: " + string(jsonPayload))
	fmt.Println("Signature: " + signature)

	respMomo, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(respMomo.Body).Decode(&result)
	fmt.Println("Response from Momo: ", result)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	return
}
