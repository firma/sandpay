package reponse

import "encoding/json"

type OrderPayResponse struct {
	Head header `json:"head"`
	Body struct {
		TotalAmount    string `json:"totalAmount"`
		ClearDate      string `json:"clearDate"`
		Credential     string `json:"credential"`
		TradeNo        string `json:"tradeNo"`
		PayTime        string `json:"payTime"`
		BuyerPayAmount string `json:"buyerPayAmount"`
		OrderCode      string `json:"orderCode"`
		DiscAmount     string `json:"discAmount"`
	} `json:"body"`
}

func (resp *OrderPayResponse) SetData(data string) {
	dataByte := []byte(data)
	json.Unmarshal(dataByte, resp)
}
