package reponse

import "encoding/json"

type OrderQueryResponse struct {
	Head header `json:"head"`
	Body struct {
		OriOrderCode   string `json:"oriOrderCode"`
		OriRespCode    string `json:"oriRespCode"`
		OriRespMsg     string `json:"oriRespMsg"`
		TotalAmount    string `json:"totalAmount"`
		OrderStatus    string `json:"orderStatus"`
		BuyerPayAmount string `json:"buyerPayAmount"`
		DiscAmount     string `json:"discAmount"`
		PayTime        string `json:"payTime"`
		ClearDate      string `json:"clearDate"`
		MidFee         string `json:"midFee"`
		PlMidFee       string `json:"plMidFee"`
		SpecialFee     string `json:"specialFee"`
		ExtraFee       string `json:"extraFee"`
		Extend         string `json:"extend"`
	} `json:"body"`
}

func (resp *OrderQueryResponse) SetData(data string) {
	dataByte := []byte(data)
	json.Unmarshal(dataByte, resp)
}
