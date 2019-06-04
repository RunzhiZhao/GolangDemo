package service

import (
	"fmt"
	"github.com/rs/xid"
	"golang_demo/db"
	"golang_demo/model"
)



// 1. 创建插入demo_order
func CreateOrderService(req *model.CreateOrderReq) (err error)  {

	// 生成随机字符串模拟orderId
	guid := xid.New()

	order := model.DemoOrder{OrderId: guid.String(), UserName: req.UserName, Amount: req.Amount, Status: "0", FileUrl: req.FileUrl}

	gormDb := db.GetDb()
	//defer  DB.Close()
	err = gormDb.Create(&order).Error
	return err
}


// 2. 更新 demo_order （amount、status、file_url）
func UpdateOrderService(req *model.UpdateOrderReq) (err error) {
	gormDb := db.GetDb()
	var order model.DemoOrder
	err = gormDb.Model(&order).Where("order_id = ?", req.OrderId).Updates(model.DemoOrder{Amount: req.Amount, Status: req.Status, FileUrl: req.FileUrl}).Error
	return err
}


// 3. 获取 demo_order 详情
func GetOrderInfoService(req *model.GetOrderInfoReq) *model.DemoOrder {
	gormDb := db.GetDb()
	var order model.DemoOrder
	if err := gormDb.Where("order_id = ?", req.OrderId).First(&order).Error; err != nil {
		return nil
	}
	return &order
}


// 4. 获取 demo_order 列表 （需要包含： 模糊查找、根据创建时间，金额排序）
func GetOrdersService(req *model.GetOrdersReq) *[]model.DemoOrder {
	gormDb := db.GetDb()
	var orders = []model.DemoOrder{}
	if err := gormDb.Where("user_name LIKE ?", req.Keyword).Find(&orders).Error; err != nil {
		fmt.Println(err)
		return nil
	}
	return &orders
}