package service

import (
	"golang_demo/db"
	"golang_demo/model"
	"testing"
)

func TestCreateOrderService(t *testing.T) {

	db.GetDb().AutoMigrate(&model.DemoOrder{})

	if err := CreateOrderService(&model.CreateOrderReq{
		UserName:"Aug",
		Amount:"998",
		FileUrl:"www.baidu.com",
	});err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateOrderService(t *testing.T) {

	db.GetDb().AutoMigrate(&model.DemoOrder{})

	if err := UpdateOrderService(&model.UpdateOrderReq{
		OrderId: "bjr1pui3q56092ji6bv0",
		Amount: "55",
		FileUrl: "www.baidu.com",
		Status: "1",
	}); err != nil {
		t.Error(err.Error())
	}
}

func TestGetOrderInfoService(t *testing.T) {

	db.GetDb().AutoMigrate(&model.DemoOrder{})

	if order := GetOrderInfoService(&model.GetOrderInfoReq{
		OrderId: "bjr1pui3q56092ji6bv0",
	}); order.OrderId!="bjr1pui3q56092ji6bv0" {
		t.Error("fail")
	}
}

func TestGetOrdersService(t *testing.T) {

	db.GetDb().AutoMigrate(&model.DemoOrder{})

	if orders := GetOrdersService(&model.GetOrdersReq{
		Keyword: "%a%",
	}); orders == nil {
		t.Error("fail")
	}
}