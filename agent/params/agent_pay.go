package params

type AgentPayParams struct {
	//Version      string `json:"version"`
	//ProductId    string `json:"productId"`
	TimeOut   string `json:"timeOut"`
	TranTime  string `json:"tranTime"`
	OrderCode string `json:"orderCode"`
	TranAmt   string `json:"tranAmt"`
	//CurrencyCode string `json:"currencyCode"`
	AccAttr  string `json:"accAttr"`
	AccType  string `json:"accType"`
	AccNo    string `json:"accNo"`
	AccName  string `json:"accName"`
	BankName string `json:"bankName"`
	Remark   string `json:"remark"`
	PayMode  string `json:"payMode"`
	//渠道类型 M 07:互联网 08:移动端 //表示商户业务使用场 //景,可默认填 07
	ChannelType string `json:"channelType"`
}

//func (params *AgentPayParams) SetVersion(object string) *AgentPayParams {
//	params.Version = object
//	return params
//}
//func (params *AgentPayParams) SetProductId(object string) *AgentPayParams {
//	params.ProductId = object
//	return params
//}
func (params *AgentPayParams) SetTranTime(object string) *AgentPayParams {
	params.TranTime = object
	return params
}
func (params *AgentPayParams) SetOrderCode(object string) *AgentPayParams {
	params.OrderCode = object
	return params
}
func (params *AgentPayParams) SetTranAmt(object string) *AgentPayParams {
	params.TranAmt = object
	return params
}

//func (params *AgentPayParams) SetCurrencyCode(object string) *AgentPayParams {
//	params.CurrencyCode = object
//	return params
//}
func (params *AgentPayParams) SetAccAttr(object string) *AgentPayParams {
	params.AccAttr = object
	return params
}
func (params *AgentPayParams) SetAccType(object string) *AgentPayParams {
	params.AccType = object
	return params
}
func (params *AgentPayParams) SetAccNo(object string) *AgentPayParams {
	params.AccNo = object
	return params
}
func (params *AgentPayParams) SetAccName(object string) *AgentPayParams {
	params.AccName = object
	return params
}
func (params *AgentPayParams) SetRemark(object string) *AgentPayParams {
	params.Remark = object
	return params
}
func (params *AgentPayParams) SetBankName(object string) *AgentPayParams {
	params.BankName = object
	return params
}
func (params *AgentPayParams) SetChannelType(object string) *AgentPayParams {
	params.ChannelType = object
	return params
}
func (params *AgentPayParams) SetPayMode(object string) *AgentPayParams {
	params.PayMode = object
	return params
}
