package model

import "github.com/jinzhu/gorm"

//用户实体
type User struct {
	gorm.Model

	UserName string `json:"user_name" gorm:"comment:'用户名'"`
	Password string `json:"password" gorm:"comment:'密码'"'`

	//UserInfo  UserInfo
	//Adresses  []Adress
	//Categorys []Category
	//Cart      Cart
	//Orders    []Orders
}
