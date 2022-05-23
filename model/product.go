package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model

	ProductName string `json:"name" gorm:"not null;unique;comment:'用户名'"`
	Img         string `json:"img" gorm:"comment:'图片'"`
	Msg         string `json:"msg" gorm:"comment:'信息'"`
	CategoryID  string `json:"category_id" gorm:"commnet:'类别ID'"`
	Price       string `json:"price" gorm:"comment:'价格'"`

	Category []Category `gorm:"many2many:category_product""`
	Orders   []Orders   `gorm:"many2many:product_orders"`
	Cart     []Cart     `gorm:"many2many:cart_product"`
}
