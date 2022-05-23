package model

import "github.com/jinzhu/gorm"

type Adress struct {
	gorm.Model

	Province string `json:"province" gorm:" comment:'省'"`
	City     string `json:"city" gorm:"comment:'市'"`
	UserID   string `json:"user_id" gorm:"comment:'用户ID'"`
	Name     string `json:"name" gorm:"comment:'姓名'"`
}
