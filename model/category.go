package model

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model

	Name    string    `json:"name" gorm:"comment:'类型名'""`
	Product []Product `gorm:"many2many:category_product;"`
}
