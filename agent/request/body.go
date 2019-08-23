package request

//整体报文格式
type AgentBody struct {
	// 交易码
	TransCode string `json:"transCode"`
	// 接入类型
	AccessType string `json:"accessType"`
	// 合作商户ID	杉德系统分配，唯一标识
	MerId string `json:"merId"`
	// 平台商户ID	平台接入必填，商户接入为空
	PlId string `json:"plId"`
	// 加密后的AES秘钥
	EncryptKey string `json:"encryptKey"`
	// 加密后的请求/应答报文
	EncryptData string `json:"encryptData"`
	// 签名
	Sign string `json:"sign"`
	// 扩展域
	Extend string `json:"extend"`
}
