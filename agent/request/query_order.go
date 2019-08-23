package request

type QueryOrderBody struct {
	Version   string `json:"version"`
	ProductId string `json:"productId"`
	TranTime  string `json:"tranTime"`
	OrderCode string `json:"orderCode"`
	Extend    string `json:"extend,omitempty"`
}
