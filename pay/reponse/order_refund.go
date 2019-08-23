package reponse

import "encoding/json"

type OrderRefundResponse struct {
	Head header `json:"head"`
	Body struct {
		Mid                 string `json:"mid"`
		OrderCode           string `json:"orderCode"`
		TotalAmount         string `json:"totalAmount"`
		OrderStatus         string `json:"orderStatus"`
		TradeNo             string `json:"tradeNo"`
		SettleAmount        string `json:"settleAmount"`
		BuyerPayAmount      string `json:"buyerPayAmount"`
		DiscAmount          string `json:"discAmount"`
		TxnCompleteTime     string `json:"txnCompleteTime"`
		PayTime             string `json:"payTime"`
		ClearDate           string `json:"clearDate"`
		AccNo               string `json:"accNo"`
		MidFee              string `json:"midFee"`
		ExtraFee            string `json:"extraFee"`
		SpecialFee          string `json:"specialFee"`
		PlMidFee            string `json:"plMidFee"`
		RefundAmount        string `json:"refundAmount"`
		SurplusAmount       string `json:"surplusAmount"`
		Bankserial          string `json:"bankserial"`
		ExternalProductCode string `json:"externalProductCode"`
		CardNo              string `json:"cardNo"`
		CreditFlag          string `json:"creditFlag"`
		Bid                 string `json:"bid"`
		Extend              string `json:"extend"`
	} `json:"body"`
}

func (resp *OrderRefundResponse) SetData(data string) {
	dataByte := []byte(data)
	json.Unmarshal(dataByte, resp)
}
