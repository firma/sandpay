package reponse

import (
	"encoding/json"
	"log"
)

type QueryOrderResponse struct {
	OrigRespCode string `json:"origRespCode"`
	OrigRespDesc string `json:"origRespDesc"`
	OrderCode    string `json:"orderCode"`
	RespCode     string `json:"respCode"`
	RespDesc     string `json:"respDesc"`
	TranTime     string `json:"tranTime"`
	ResultFlag   string `json:"resultFlag"`
	SandSerial   string `json:"sandSerial"`
	TranDate     string `json:"tranDate"`
	TranFee      string `json:"tranFee"`
	ExtraFee     string `json:"extraFee"`
	HolidayFee   string `json:"holidayFee"`
	Extend       string `json:"extend"`
}

func (resp *QueryOrderResponse) SetData(data string) {
	json.Unmarshal([]byte(data), resp)
	log.Println(resp, data)
}
