package params

import "fmt"

//支付参数定义
type OrderRefundParams struct {
	//商户订单号
	OrderNo string `json:"orderCode"`
	//2. 原商户订单号
	RefundNO string `json:"oriOrderCode"`
	//3. 退款金额
	RefundAmount int
	//10. 退货原因
	RefundReason string `json:"refundReason"`
	//19. 扩展域
	Extends string `json:"extends"`
}

func (params *OrderRefundParams) SetOrderNo(orderNo string) *OrderRefundParams {
	params.OrderNo = orderNo
	return params
}
func (params *OrderRefundParams) setRefundNo(orderNo string) *OrderRefundParams {
	params.RefundNO = orderNo
	return params
}

func (params *OrderRefundParams) SetRefundAmount(object int) *OrderRefundParams {
	params.RefundAmount = object
	return params
}
func (params *OrderRefundParams) SetRefundReason(object string) *OrderRefundParams {
	params.RefundReason = object
	return params
}
func (params *OrderRefundParams) SetExtends(object string) *OrderRefundParams {
	params.Extends = object
	return params
}

func (params *OrderRefundParams) GetRefundAmount() string {
	amount := fmt.Sprintf("%012d", params.RefundAmount)
	return amount
}
