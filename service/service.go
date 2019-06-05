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
	if err = gormDb.Create(&order).Error; err != nil {
		return err
	}
	return nil
}

// 事务
func CreateOrderService1(req *model.CreateOrderReq) (err error)  {

	// 生成随机字符串模拟orderId
	guid := xid.New()

	order := model.DemoOrder{OrderId: guid.String(), UserName: req.UserName, Amount: req.Amount, Status: "0", FileUrl: req.FileUrl}

	gormDb := db.GetDb()

	gormDb.Begin()

	if err = gormDb.Create(&order).Error; err != nil {
		gormDb.Rollback()
		return err
	}

	gormDb.CommonDB()

	return nil
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

	var page uint = 1
	if req.Page > 1 {
		page = req.Page
	}

	var pageSize uint = 5
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}

	if len(req.Keyword) > 0 {	// 模糊查找
		if err := gormDb.Where("user_name LIKE ?", req.Keyword).Offset((page - 1) * pageSize).Limit(pageSize).Find(&orders).Error; err != nil {
			panic(err.Error())
			return nil
		}
	} else if req.SortType == 1 {	// 根据创建时间排序

		if err := gormDb.Order("created_at").Offset((page - 1) * pageSize).Limit(pageSize).Find(&orders).Error; err != nil {
			panic(err.Error())
			return nil
		}

	} else if req.SortType == 2 {	// 根据金额排序

		if err := gormDb.Order("amount").Offset((page - 1) * pageSize).Limit(pageSize).Find(&orders).Error; err != nil {
			panic(err.Error())
			return nil
		}
	}

	fmt.Println(len(orders))
	fmt.Println("..\n..")
	fmt.Print(orders)

	return &orders
}