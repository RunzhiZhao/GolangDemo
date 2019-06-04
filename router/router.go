package router

import (
	"github.com/gin-gonic/gin"
	"golang_demo/handler"
)

func StartRouter()  {
	router := gin.Default()

	// 1. 创建插入demo_order
	router.POST("/createOrder", handler.CreateOrder)

	// 2. 更新 demo_order （amount、status、file_url）
	router.POST("/updateOrder", handler.UpdateOrder)

	// 3. 获取 demo_order 详情
	router.GET("/getOrderInfo", handler.GetOrderInfo)

	// 4. 获取 demo_order 列表 （需要包含： 模糊查找、根据创建时间，金额排序）
	router.GET("/getOrders", handler.GetOrders)

	router.Run("localhost:3306")
}
