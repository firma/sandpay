package request

type OrderQueryBody struct {
	//商户订单号
	OrderCode string `json:"orderCode"`
	//19. 扩展域
	Extends string `json:"extends"`
}
