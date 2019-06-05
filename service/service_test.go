package service

import (
	"golang_demo/db"
	"golang_demo/model"
	"testing"
)

/* 创建订单 */
func TestCreateOrderService(t *testing.T) {

	db.GetDb().AutoMigrate(&model.DemoOrder{})

	if order := CreateOrderService(&model.CreateOrderReq{
		UserName: "user04",
		Amount:   29.9,
		FileUrl:  "www.baidu.com",
	}); order == nil {
		t.Error("error")
	}
}

/* SQL事务 创建订单 */
func TestCreateOrderService1(t *testing.T) {

	db.GetDb().AutoMigrate(&model.DemoOrder{})

	if order := CreateOrderService1(&model.CreateOrderReq{
		UserName: "User07",
		Amount: 19.9,
		FileUrl: "github.com",
	}); order == nil {
		t.Error("error")
	}
}

/* 更新订单信息 */
func TestUpdateOrderService(t *testing.T) {

	db.GetDb().AutoMigrate(&model.DemoOrder{})

	if order := UpdateOrderService(&model.UpdateOrderReq{
		OrderId: "bjroh5q3q561jbctp23g",
		Amount:  79.9,
		FileUrl: "www.baidu.com",
		Status:  "1",
	}); order == nil {
		t.Error("error")
	}
}

/* 获取订单信息 */
func TestGetOrderInfoService(t *testing.T) {

	db.GetDb().AutoMigrate(&model.DemoOrder{})

	if order := GetOrderInfoService(&model.GetOrderInfoReq{
		OrderId: "bjr3h8i3q560itdtgjp0",
	}); order.OrderId != "bjr3h8i3q560itdtgjp0" {
		t.Error("fail")
	}
}

/* 模糊查找 */
func TestGetOrdersService(t *testing.T) {

	db.GetDb().AutoMigrate(&model.DemoOrder{})

	if orders := GetOrdersService(&model.GetOrdersReq{
		Keyword: "a",
	}); orders == nil {
		t.Error("fail")
	}
}

/* 按创建时间排序 */
func TestGetOrdersService2(t *testing.T) {
	db.GetDb().AutoMigrate(&model.DemoOrder{})

	if orders := GetOrdersService(&model.GetOrdersReq{
		SortType: 1,
		Page:     1,
		PageSize: 3,
	}); orders == nil {
		t.Error("fail")
	}
}

/* 按金额小排序 */
func TestGetOrdersService3(t *testing.T) {
	db.GetDb().AutoMigrate(&model.DemoOrder{})

	if orders := GetOrdersService(&model.GetOrdersReq{
		SortType: 2,
		Page:     1,
		PageSize: 3,
	}); orders == nil {
		t.Error("fail")
	}
}

func TestDataTableToExcelService(t *testing.T) {
	db.GetDb().AutoMigrate(&model.DemoOrder{})
	if err := DataTableToExcelService(); err != nil {
		t.Error("fail")
	}
}