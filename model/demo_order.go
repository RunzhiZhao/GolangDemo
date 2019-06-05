package model

import "github.com/jinzhu/gorm"

type DemoOrder struct {
	gorm.Model
	OrderId  string`gorm:"unique;not null"`
	UserName string
	Amount   float64
	Status   string
	FileUrl  string
}
