package request

type PostData struct {
	Charset  string `json:"charset"`
	SignType string `json:"ignType"`
	Data     string `json:"ata"`
	Sign     string `json:"sign"`
}

type Header struct {
	Version    string `json:"version"`
	Method     string `json:"method"`
	Mid        string `json:"mid"`
	AccessType string `json:"accessType"`
	//PlMid       string `json:"plMid"`
	ChannelType string `json:"channelType"`
	ReqTime     string `json:"reqTime"`
	ProductId   string `json:"productId"`
}

func (h *Header) SetVersion(version string) *Header {
	h.Version = version
	return h
}
func (h *Header) SetMethod(method string) *Header {
	h.Method = method
	return h
}
func (h *Header) SetMid(mid string) *Header {
	h.Mid = mid
	return h
}

//func (h Header) SetPlMid(plMid string) {
//	h.PlMid = plMid
//}
func (h *Header) SetAccessType(accessType string) *Header {
	h.AccessType = accessType
	return h
}
func (h *Header) SetReqTime(reqTime string) *Header {
	h.ReqTime = reqTime
	return h
}
func (h *Header) SetProductId(productId string) *Header {
	h.ProductId = productId
	return h
}
func (h *Header) SetChannelType(channelType string) *Header {
	h.ChannelType = channelType
	return h
}
