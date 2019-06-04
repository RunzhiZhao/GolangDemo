package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang_demo/model"
	"sync"
)

var db *gorm.DB
var once sync.Once

func GetDb() *gorm.DB {
	once.Do(func() {

		dbConfig := MyDbConfig()
		if !checkConfig(&dbConfig) {
			panic("db配置不正确")
		}

		cStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
		var err error
		if db, err = gorm.Open("mysql", cStr); err != nil {
			panic("连接数据库出错:" + err.Error())
		}
	})

	db.AutoMigrate(&model.DemoOrder{})

	return db
}

// 创建数据库
func CreateDB(dbConfig *DbConfig) {

	if !checkConfig(dbConfig) {
		panic("db配置不正确")
	}
	cStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, "information_schema")
	openedDb, err := gorm.Open("mysql", cStr)

	if err != nil {
		fmt.Println(cStr)
		panic("连接数据库出错:" + err.Error())
	}

	createDbSQL := "CREATE DATABASE IF NOT EXISTS " + dbConfig.DbName + " DEFAULT CHARSET utf8 COLLATE utf8_general_ci;"

	err = openedDb.Exec(createDbSQL).Error
	if err != nil {
		fmt.Println("创建失败：" + err.Error() + " sql:" + createDbSQL)
		return
	}
	fmt.Println(dbConfig.DbName + "数据库创建成功")

	openedDb.AutoMigrate(&model.DemoOrder{})

	defer openedDb.Close()
}

/// 检测数据库配置
func checkConfig(dbConfig *DbConfig) bool {
	return true
}
