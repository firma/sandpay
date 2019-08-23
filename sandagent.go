package sandpay

import (
	"log"
	"sandpay/agent"
	"sandpay/agent/params"
	"sandpay/agent/reponse"
	"sandpay/agent/request"
	"time"
)

var AgentPayClient PaymentAgent

func init() {
	//TODO 改为配置文件
	AgentPayClient.Config = agent.Config{
		MerId: "100211701160001",
		PrivatePath:     "/home/keyed/Downloads/100211701160001/server.key",  //私钥文件
		CertPath:        "/home/keyed/Downloads/100211701160001/server.cert", //公钥文件
		EncryptCertPath: "/home/keyed/Downloads/100211701160001/sand.cer", //导出的公钥文件

		ApiHost: "https://caspay.sandpay.com.cn/agent-main/openapi",

		NotifyUrl: "https://61.129.71.103:8003/pay/api/",
	}
	agent.LoadCertInfo(&AgentPayClient.Config)
}

type PaymentAgent struct {
	Config agent.Config
}

// 实时代付
func (sandPay *PaymentAgent) AgentPay(params params.AgentPayParams) (resp reponse.AgentPayResponse, err error) {
	timeString := time.Now().Format("20060102150405")

	body := request.AgentPayBody{
		Version:      agent.VERSION,
		OrderCode:    params.OrderCode,
		ProductId:    agent.PRODUCTID_AGENTPAY_TOC,
		CurrencyCode: agent.CURRENCY_CODE,
		TranTime:     timeString,
		TimeOut:      params.TimeOut,
		TranAmt:      params.TranAmt,
		AccAttr:      params.AccAttr,
		AccNo:        params.AccNo,
		AccType:      params.AccType,
		AccName:      params.AccName,
		BankName:     params.BankName,
		Remark:       params.Remark,
		PayMode:      params.PayMode,
		ChannelType:  params.ChannelType,
	}

	signDataJsonString := agent.GenerateSignString(body)
	//log.Println("请求提中 encryptData 原始数据  ", signDataJsonString)
	encryptData, sign, encryptKey, err := agent.PrivateSha1SignData(signDataJsonString)
	postData := agent.GeneratePostData(encryptData, encryptKey, agent.AGENT_PAY, sandPay.Config.MerId, sign)
	//log.Println("发送的请求体body内容", postData)
	data, err := agent.PayPost(sandPay.Config.ApiHost+"/agentpay", postData)
	//fmt.Println("过滤 \" 解析内容", data, err)
	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}

// 订单查询
func (sandPay *PaymentAgent) AgentPayQuery(tranTime string, orderCode string) (resp reponse.QueryOrderResponse, err error) {

	body := request.QueryOrderBody{
		Version:   agent.VERSION,
		ProductId: agent.PRODUCTID_AGENTPAY_TOC,
		TranTime:  tranTime,
		OrderCode: orderCode,
		Extend:    "",
	}

	signDataJsonString := agent.GenerateSignString(body)
	log.Println(signDataJsonString)
	encryptData, sign, encryptKey, err := agent.PrivateSha1SignData(signDataJsonString)
	postData := agent.GeneratePostData(encryptData, encryptKey, agent.ORDER_QUERY, sandPay.Config.MerId, sign)

	data, err := agent.PayPost(sandPay.Config.ApiHost+"/queryOrder", postData)
	if err != nil {
		return
	}
	resp.SetData(data.Data)
	return resp, err
}
