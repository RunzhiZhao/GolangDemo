package model

// /createOrder接口请求参数
type CreateOrderReq struct {
	UserName string  `json:"user_name"`
	Amount   float64 `json:"amount"`
	FileUrl  string  `json:"file_url"`
}

// /updateOrder接口请求参数
type UpdateOrderReq struct {
	OrderId string  `json:"order_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
	FileUrl string  `json:"file_url"`
}

// /getOrderInfo接口请求参数
type GetOrderInfoReq struct {
	OrderId string `json:"order_id"`
}

// /getOrders接口请求参数, 需要包含: 模糊查找、根据创建时间，金额排序）
type GetOrdersReq struct {
	Keyword  string `json:"keyword"`   /* 关键字模糊查找 */
	SortType int    `json:"sort_type"` /* 1: 根据创建时间排序， 2根据金额排序*/
	Page     uint   `json:"page"`      /* 页码 */
	PageSize uint   `json:"page_size"` /* 每页数量 */
}
