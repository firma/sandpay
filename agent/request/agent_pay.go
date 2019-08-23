package request

type AgentPayBody struct {
	Version      string `json:"version"`
	ProductId    string `json:"productId"`
	TranTime     string `json:"tranTime"`
	OrderCode    string `json:"orderCode"`
	TimeOut      string `json:"timeOut,omitempty"`
	TranAmt      string `json:"tranAmt"`
	CurrencyCode string `json:"currencyCode"`
	AccAttr      string `json:"accAttr"`
	AccType      string `json:"accType"`
	AccNo        string `json:"accNo"`
	AccName      string `json:"accName"`
	ProvNo       string `json:"provNo,omitempty"`
	CityNo       string `json:"cityNo,omitempty"`
	BankName     string `json:"bankName"`
	BankType     string `json:"bankType,omitempty"`
	Remark       string `json:"remark"`
	PayMode      string `json:"payMode"`
	//渠道类型 M 07:互联网 08:移动端 //表示商户业务使用场 //景,可默认填 07
	ChannelType string `json:"channelType"`
	//业务扩展参数 JSON 结构{ busiType:业务类型 2-提现 storeId:门店编号 ANS0..20 accountDate:商户记 账日期 ANS8}
	ExtendParams string `json:"extendParams,omitempty"`
	ReqReserved  string `json:"reqReserved,omitempty"`
	NoticeUrl    string `json:"noticeUrl,omitempty"`
	Extend       string `json:"extend,omitempty"`
	Phone        string `json:"phone,omitempty"`
	//1:同城 2:异地
	AreaFlag string `json:"areaFlag,omitempty"`
	//银行机构编码
	ChannelBankCode string `json:"channelBankCode,omitempty"`
	//01-资金结算 //02-代付 //03-支付账户提现 //04-退单重划 //05-预付费卡-售卡款赎 //06-资金调拨 //07-资金结算 //08-结转利息收入 09-结转手续费收入 10-专户回款
	SettlementType string `json:"settlement_type,omitempty"`
}
