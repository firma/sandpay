package reponse

import "encoding/json"

type AgentPayResponse struct {
	RespCode   string `json:"respCode"`
	RespDesc   string `json:"respDesc"`
	TranTime   string `json:"tranTime"`
	OrderCode  string `json:"orderCode"`
	ResultFlag string `json:"resultFlag"`
	SandSerial string `json:"sandSerial"`
	TranDate   string `json:"tranDate"`
	TranFee    string `json:"tranFee"`
	ExtraFee   string `json:"extraFee"`
	HolidayFee string `json:"holidayFee"`
	Extend     string `json:"extend"`
}

func (resp *AgentPayResponse) SetData(data string) {
	dataByte := []byte(data)
	json.Unmarshal(dataByte, resp)
}
