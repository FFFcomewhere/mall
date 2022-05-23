package user_service

import (
	"exam8/global"
	"exam8/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

func ExistUserByUserName(u model.User) (bool, error) {
	if err := global.DB.Where("user_name = ?", u.UserName).First(&model.User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}
	return true, nil
}

func AddUser(u model.User) (*model.User, error) {
	fmt.Println("开始添加用户")
	fmt.Println("u", u)

	if err := global.DB.Create(&u).Error; err != nil {
		return nil, err
	}

	//if err := global.DB.Where("user_name = ?", u.UserInfo.UserName).First(&u).Error; err != nil {
	if err := global.DB.Where("user_name = ?", u.UserName).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func GetUser(u model.User) (*model.User, error) {
	if err := global.DB.Where("user_name = ?", u.UserName).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func UpdateUser(u model.User) (*model.User, error) {
	if err := global.DB.Model(model.User{}).Where("user_name = ?", u.UserName).Update(u).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("user_name = ?", u.UserName).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}

func DeleteUser(u model.User) error {
	fmt.Println(u.UserName)
	if err := global.DB.Delete(model.User{}, "user_name = ?", u.UserName).Error; err != nil {
		return err
	}

	return nil
}

func UserLogin(u model.User) (*model.User, error) {
	var newUser model.User

	if err := global.DB.Where("user_name = ?", u.UserName).First(&newUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	if u.Password != u.Password {
		return nil, nil
	}

	return &newUser, nil
}
