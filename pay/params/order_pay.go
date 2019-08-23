package params

import (
	"encoding/json"
	"fmt"
)

const PayModWeiXinMp = "sand_wx"

//支付参数定义
type OrderPayParams struct {
	//商户订单号
	OrderNo string
	//2. 订单金额
	TotalAmount int
	//3. 订单标题
	Subject string
	//4. 订单描述
	Body string
	//5. 订单超时时间
	TxnTimeOut string
	//	8. 客户端 IP
	ClientIp string
	//支付模式
	PayMode string `json:"payMode"`
	//	7. 支付扩展域  ANS0.1024 C 具体格式根据 payMode 确定,
	PayExtra PayExtraWeiChat `json:"payExtra"`
	//19. 扩展域
	Extends string
}

func (params *OrderPayParams) SetPayMode(object string) *OrderPayParams {
	params.PayMode = object
	return params
}
func (params *OrderPayParams) SetPayExtra(openId string, AppId string) *OrderPayParams {
	params.PayExtra = PayExtraWeiChat{
		SubAppId: AppId,
		OpenId:   openId,
	}
	return params
}

func (params *OrderPayParams) SetOrderNo(orderNo string) *OrderPayParams {
	params.OrderNo = orderNo
	return params
}
func (params *OrderPayParams) SetTotalAmount(object int) *OrderPayParams {
	params.TotalAmount = object
	return params
}

func (params *OrderPayParams) GetTotalAmountToString() string {
	amount := fmt.Sprintf("%012d", params.TotalAmount)
	return amount
}

func (params *OrderPayParams) SetSubject(object string) *OrderPayParams {
	params.Subject = object
	return params
}
func (params *OrderPayParams) SetBody(object string) *OrderPayParams {
	params.Body = object
	return params
}
func (params *OrderPayParams) SetTxnTimeOut(object string) *OrderPayParams {
	params.TxnTimeOut = object
	return params
}
func (params *OrderPayParams) SetClientIp(object string) *OrderPayParams {
	params.ClientIp = object
	return params
}

type PayExtraWeiChat struct {
	SubAppId string `json:"subAppid"`
	OpenId   string `json:"userId"`
}

func (extra *PayExtraWeiChat) SetOpenId(object string) *PayExtraWeiChat {
	extra.OpenId = object
	return extra
}
func (extra *PayExtraWeiChat) SetSubAppId(object string) *PayExtraWeiChat {
	extra.SubAppId = object
	return extra
}
func (extra *PayExtraWeiChat) ToJson() string {
	extraByte, err := json.Marshal(extra)
	if err != nil {
		return ""
	}
	return string(extraByte[:])
}
