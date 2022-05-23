package model

import "github.com/jinzhu/gorm"

type UserInfo struct {
	gorm.Model

	UserName string `json:"user_name" gorm:"comment:'用户名'"`
	Password string `json:"password" gorm:"comment:'密码'"'`
	Sex      string `json:"gender" gorm:"comment:'性别'"`
	PhoneNum string `json:"phone_num" gorm:"commnet:'手机号码'"`
	Score    int    `json:"score" gorm:"comment:'积分'"'`
}
