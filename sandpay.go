package sandpay

import (
	"net/url"
	"sandpay/pay"
	"sandpay/pay/params"
	"sandpay/pay/reponse"
	"sandpay/pay/request"
	"strings"
	"time"
)

var Client SandPay

func init() {
	//TODO 改为配置文件
	Client.Config = pay.Config{
		MerId: "100211701160001",
		PrivatePath:     "/home/keyed/Downloads/100211701160001/server.key",  //私钥文件
		CertPath:        "/home/keyed/Downloads/100211701160001/server.cert", //公钥文件
		EncryptCertPath: "/home/keyed/Downloads/100211701160001/sand.cer", //导出的公钥文件


		ApiHost:   "https://cashier.sandpay.com.cn/gateway/api",
		NotifyUrl: "https://127.0.0.1:8080/pay/api/",
	}
	pay.LoadCertInfo(&Client.Config)
}

type SandPay struct {
	Config pay.Config
}

// 统一下单接口
func (sandPay *SandPay) OrderPay(params params.OrderPayParams) (resp reponse.OrderPayResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.pay`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00000005").SetReqTime(timeString)
	body := request.OrderPayBody{
		OrderCode:   params.OrderNo,
		TotalAmount: params.GetTotalAmountToString(),
		Subject:     params.Subject,
		Body:        params.Body,
		TxnTimeOut:  params.TxnTimeOut,
		PayMode:     params.PayMode,
		PayExtra:    params.PayExtra.ToJson(),
		ClientIp:    params.ClientIp,
		NotifyUrl:   sandPay.Config.NotifyUrl,
		FrontUrl:    sandPay.Config.FrontUrl,
		Extends:     params.Extends,
	}

	signDataJsonString := pay.GenerateSignString(body, header)
	sign, _ := pay.PrivateSha1SignData(signDataJsonString)
	postData := pay.GeneratePostData(signDataJsonString, sign)

	data, err := pay.PayPost(config.ApiHost+"/order/pay", postData)
	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}

// 订单查询接口
func (sandPay *SandPay) OrderQuery(orderNo string, extend string) (resp reponse.OrderQueryResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.query`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00000007").SetReqTime(timeString)
	body := request.OrderQueryBody{
		OrderCode: orderNo,
		Extends:   extend,
	}

	signDataJsonString := pay.GenerateSignString(body, header)
	sign, _ := pay.PrivateSha1SignData(signDataJsonString)
	postData := pay.GeneratePostData(signDataJsonString, sign)

	data, err := pay.PayPost(config.ApiHost+"/order/query", postData)

	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}

// 退货申请接口
func (sandPay *SandPay) OrderRefund(refundParams params.OrderRefundParams) (resp reponse.OrderRefundResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.refund`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00000007").SetReqTime(timeString)
	body := request.OrderRefundBody{
		OrderCode:    refundParams.OrderNo,
		OriOrderCode: refundParams.RefundNO,
		RefundAmount: refundParams.GetRefundAmount(),
		NotifyUrl:    config.NotifyUrl,
		RefundReason: refundParams.RefundReason,
		Extends:      refundParams.Extends,
	}

	signDataJsonString := pay.GenerateSignString(body, header)
	sign, _ := pay.PrivateSha1SignData(signDataJsonString)
	postData := pay.GeneratePostData(signDataJsonString, sign)

	data, err := pay.PayPost(config.ApiHost+"/order/refund", postData)
	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}

// 退货申请接口
func (sandPay *SandPay) OrderRefunds(refundParams params.OrderRefundParams) (resp reponse.OrderRefundResponse, err error) {
	config := sandPay.Config
	timeString := time.Now().Format("20060102150405")

	header := request.Header{}
	header.SetMethod(`sandpay.trade.refund`).SetVersion(`1.0`).SetAccessType("1")
	header.SetChannelType("07").SetMid(config.MerId).SetProductId("00000007").SetReqTime(timeString)
	body := request.OrderRefundBody{
		OrderCode:    refundParams.OrderNo,
		OriOrderCode: refundParams.RefundNO,
		RefundAmount: refundParams.GetRefundAmount(),
		NotifyUrl:    config.NotifyUrl,
		RefundReason: refundParams.RefundReason,
		Extends:      refundParams.Extends,
	}

	signDataJsonString := pay.GenerateSignString(body, header)
	sign, _ := pay.PrivateSha1SignData(signDataJsonString)
	postData := pay.GeneratePostData(signDataJsonString, sign)

	data, err := pay.PayPost(config.ApiHost+"/order/refund", postData)
	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}

// queryString 回调参数校验
func NotifyVerifyData(dataString string) (response reponse.Response, err error) {
	var fields []string
	fields = strings.Split(dataString, "&")

	vals := url.Values{}
	for _, field := range fields {
		f := strings.SplitN(field, "=", 2)
		if len(f) >= 2 {
			key, val := f[0], f[1]
			vals.Set(key, val)
		}
	}

	result, err := pay.PublicSha1Verify(vals)

	if err != nil {
		return response, err
	}
	mapInfo := result.(map[string]string)
	for key, value := range mapInfo {
		response.SetKeyValue(key, value)
	}
	return response, err
}
