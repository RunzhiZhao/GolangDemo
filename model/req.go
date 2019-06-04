package model

// /createOrder接口请求参数
type CreateOrderReq struct {
	UserName string
	Amount   string
	FileUrl  string
}

// /updateOrder接口请求参数
type UpdateOrderReq struct {
	OrderId string`gorm:"unique;not null"`
	Amount  string
	Status  string
	FileUrl string
}

// /getOrderInfo接口请求参数
type GetOrderInfoReq struct {
	OrderId string
}

// /getOrders接口请求参数, 需要包含: 模糊查找、根据创建时间，金额排序）
type GetOrdersReq struct {
	Keyword string
}
