package model

import "github.com/jinzhu/gorm"

type DemoOrder struct {
	gorm.Model
	OrderId  string`gorm:"unique;not null"`
	UserName string
	Amount   string
	Status   string
	FileUrl  string
}
