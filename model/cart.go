package model

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model

	UserID  string    `json:"user_id" gorm:"comment:'用户ID'"`
	Product []Product `gorm:"many2many:cart_product"`
	Orders  []Orders
}
