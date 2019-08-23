package request

type OrderPayBody struct {
	//商户订单号
	OrderCode string `json:"orderCode"`
	//2. 订单金额
	TotalAmount string `json:"totalAmount"`
	//3. 订单标题
	Subject string `json:"subject"`
	//4. 订单描述
	Body string `json:"body"`
	//5. 订单超时时间
	TxnTimeOut string `json:"txnTimeOut"`
	//支付模式
	PayMode string `json:"payMode"`
	//	7. 支付扩展域  ANS0.1024 C 具体格式根据 payMode 确定,
	//PayExtra PayExtra `json:"payExtra"`
	PayExtra string `json:"payExtra"`
	//	8. 客户端 IP
	ClientIp string `json:"clientIp"`
	//9. 异步通知地址 notifyUrl ANS0.256 M \
	NotifyUrl string `json:"notifyUrl"`
	//10. 前台通知地址
	FrontUrl string `json:"frontUrl"`
	//ANS0.256 M 11. 商户门店编号
	StoreId string `json:"storeId"`
	//12. 商户终端编号
	TerminalId string `json:"terminalId"`
	//13. 操作员编号
	OperatorId string `json:"operatorId"`
	//14. 清算模式
	ClearCycle int `json:"clearCycle"`
	//	分账信息
	RoyaltyInfo string `json:"royaltyInfo"`
	//	16. 风控信息域
	RiskRateInfo string `json:"riskRateInfo"`
	//	17. 业务扩展参数
	BizExtendParams string `json:"bizExtendParams"`
	//	18. 商户扩展参数
	MerchExtendParams string `json:"merchExtendParams"`
	//19. 扩展域
	Extends string `json:"extends"`
}
