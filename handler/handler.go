package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_demo/model"
	"golang_demo/service"
	"net/http"
)

// 1. 创建插入demo_order
func CreateOrder(c *gin.Context) {

	var req model.CreateOrderReq
	// 参数格式检测
	if err := c.BindJSON(&req); err != nil {
		response := model.ResponseData{Code: 0, Status: "error", Message: "参数错误", Data: ""}
		c.JSON(http.StatusOK, response)
		return
	}

	if order := service.CreateOrderService(&req); order == nil {
		response := model.ResponseData{Code: 0, Status: "error", Message: "订单创建失败", Data: ""}
		c.JSON(http.StatusOK, response)
		return
	} else {
		response := model.ResponseData{Code: 1, Status: "success", Message: "操作成功", Data: order}
		c.JSON(http.StatusOK, response)
	}
}

// 2. 更新 demo_order （amount、status、file_url）
func UpdateOrder(c *gin.Context) {
	var req model.UpdateOrderReq
	if err := c.BindJSON(&req); err != nil {
		response := model.ResponseData{Code: 0, Status: "error", Message: "参数错误", Data: ""}
		c.JSON(http.StatusOK, response)
		return
	}

	if len(req.OrderId) == 0 {
		response := model.ResponseData{Code: 0, Status: "error", Message: "缺少参数order_id", Data: ""}
		c.JSON(http.StatusOK, response)
		return
	}

	if order := service.UpdateOrderService(&req); order == nil {
		response := model.ResponseData{Code: 0, Status: "error", Message: "操作失败", Data: ""}
		c.JSON(http.StatusOK, response)
		return
	} else {
		response := model.ResponseData{Code: 1, Status: "success", Message: "操作成功", Data: order}
		c.JSON(http.StatusOK, response)
	}
}

// 3. 获取 demo_order 详情
func GetOrderInfo(c *gin.Context) {

	orderId := c.Query("order_id")
	fmt.Println("order_id = " + orderId)

	if len(orderId) == 0 {
		response := model.ResponseData{Code: 0, Status: "error", Message: "缺少参数order_id", Data: ""}
		c.JSON(http.StatusOK, response)
		return
	}

	var req = model.GetOrderInfoReq{OrderId: orderId}

	if order := service.GetOrderInfoService(&req); order != nil {
		response := model.ResponseData{Code: 1, Status: "success", Message: "操作成功", Data: order}
		c.JSON(http.StatusOK, response)
		return
	} else {
		response := model.ResponseData{Code: 0, Status: "error", Message: "订单不存在", Data: ""}
		c.JSON(http.StatusOK, response)
	}
}

// 4. 获取 demo_order 列表 （需要包含： 模糊查找、根据创建时间，金额排序）
func GetOrders(c *gin.Context) {

	var req model.GetOrdersReq

	if err := c.BindJSON(&req); err != nil {
		response := model.ResponseData{Code: 0, Status: "error", Message: "参数错误", Data: ""}
		c.JSON(http.StatusOK, response)
		return
	}

	if orders := service.GetOrdersService(&req); orders != nil {
		response := model.ResponseData{Code: 1, Status: "success", Message: "操作成功", Data: orders}
		c.JSON(http.StatusOK, response)
	} else {
		response := model.ResponseData{Code: 0, Status: "error", Message: "操作失败", Data: ""}
		c.JSON(http.StatusOK, response)
	}

}


// 5. Gin 文件上传的功能， 并且更新 demo_order ： file_url
func UploadFile(c *gin.Context)  {

	// Source
	file, err := c.FormFile("file")

	if err != nil {
		response := model.ResponseData{Code: 0, Status: "error", Message: "获取文件失败", Data: ""}
		c.JSON(http.StatusOK, response)
		return
	}

	filename := "./file/" + file.Filename

	if err := c.SaveUploadedFile(file, filename); err != nil {
		response := model.ResponseData{Code: 0, Status: "error", Message: "保存文件出错", Data: ""}
		c.JSON(http.StatusOK, response)
		return
	}

	response := model.ResponseData{Code: 1, Status: "success", Message: "上传成功", Data: ""}
	c.JSON(http.StatusOK, response)
}


// 6. Gin 文件下载
func DownloadFile(c *gin.Context)  {

	filename := "huge.gpeg"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./file/gpeg")
}


// 7. 将demo_order 所有数据以excel形式导出来
func ExportOrderExcel(c *gin.Context)  {

	if err := service.DataTableToExcelService(); err != nil {
		response := model.ResponseData{Code: 0, Status: "error", Message: err.Error(), Data: ""}
		c.JSON(http.StatusOK, response)
		return
	}

	filename := "order_demo.xlsx"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("order_demo.xlsx")

	response := model.ResponseData{Code: 0, Status: "success", Message: "下载order_excel.xlsx", Data: ""}
	c.JSON(http.StatusOK, response)
}