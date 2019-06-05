package service

import (
	"fmt"
	"github.com/rs/xid"
	"github.com/tealeg/xlsx"
	"golang_demo/db"
	"golang_demo/model"
)



// 1. 创建插入demo_order
func CreateOrderService(req *model.CreateOrderReq) *model.DemoOrder  {

	// 生成随机字符串模拟orderId
	guid := xid.New()

	order := model.DemoOrder{OrderId: guid.String(), UserName: req.UserName, Amount: req.Amount, Status: "0", FileUrl: req.FileUrl}

	gormDb := db.GetDb()
	//defer  DB.Close()
	if err := gormDb.Create(&order).Error; err != nil {
		panic(err.Error())
		return nil
	}
	return &order
}

// 事务
func CreateOrderService1(req *model.CreateOrderReq) *model.DemoOrder  {

	// 生成随机字符串模拟orderId
	guid := xid.New()

	order := model.DemoOrder{OrderId: guid.String(), UserName: req.UserName, Amount: req.Amount, Status: "0", FileUrl: req.FileUrl}

	gormDb := db.GetDb()

	gormDb.Begin()

	if err := gormDb.Create(&order).Error; err != nil {
		gormDb.Rollback()
		panic(err.Error())
		return nil
	}

	gormDb.CommonDB()

	return &order
}


// 2. 更新 demo_order （amount、status、file_url）
func UpdateOrderService(req *model.UpdateOrderReq) *model.DemoOrder {
	gormDb := db.GetDb()
	var order model.DemoOrder
	if err := gormDb.Model(&order).Where("order_id = ?", req.OrderId).Updates(model.DemoOrder{Amount: req.Amount, Status: req.Status, FileUrl: req.FileUrl}).Error; err != nil {
		panic(err.Error())
		return nil
	}
	gormDb.Where("order_id = ?", req.OrderId).Find(&order)
	return &order
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
	var orders []model.DemoOrder

	var page uint = 1
	if req.Page > 1 {
		page = req.Page
	}

	var pageSize uint = 5
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}

	if len(req.Keyword) > 0 {	// 模糊查找
		keyword := "%" + req.Keyword + "%"
		if err := gormDb.Where("user_name LIKE ?", keyword).Offset((page - 1) * pageSize).Limit(pageSize).Find(&orders).Error; err != nil {
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

	return &orders
}



// 7. demo_order表转成excel
func DataTableToExcelService() (err error) {

	gormDb := db.GetDb()

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	// 创建标题
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.SetValue("id")
	cell = row.AddCell()
	cell.SetValue("created_at")
	cell = row.AddCell()
	cell.SetValue("updated_at")
	cell = row.AddCell()
	cell.SetValue("deleted_at")
	cell = row.AddCell()
	cell.SetValue("order_id")
	cell = row.AddCell()
	cell.SetValue("user_name")
	cell = row.AddCell()
	cell.SetValue("amount")
	cell = row.AddCell()
	cell.SetValue("status")
	cell = row.AddCell()
	cell.SetValue("file_url")

	// 获取数据
	var orders []model.DemoOrder
	if err = gormDb.Order("created_at").Find(&orders).Error; err != nil {
		panic(err.Error())
		return nil
	}

	for _, value := range orders {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.SetValue(value.ID)
		cell = row.AddCell()
		cell.SetValue(value.CreatedAt)
		cell = row.AddCell()
		cell.SetValue(value.UpdatedAt)
		cell = row.AddCell()
		cell.SetValue(value.DeletedAt)
		cell = row.AddCell()
		cell.SetValue(value.OrderId)
		cell = row.AddCell()
		cell.SetValue(value.UserName)
		cell = row.AddCell()
		cell.SetValue(value.Amount)
		cell = row.AddCell()
		cell.SetValue(value.Status)
		cell = row.AddCell()
		cell.SetValue(value.FileUrl)
	}

	if err = file.Save("./file/order_demo.xlsx"); err != nil {
		fmt.Printf(err.Error())
		return err
	}

	return nil

}