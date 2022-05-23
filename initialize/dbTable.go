package initialize

import (
	"exam8/global"
	"exam8/model"
)

// 注册数据库表专用
func DBTables() {
	// 因为同步最新的api接口到数据库，所以要删除旧表
	if global.DB.HasTable("apis") {
		global.DB.DropTable("apis")
	}

	global.DB.AutoMigrate(
		model.Adress{},
		model.Cart{},
		model.Category{},
		model.Orders{},
		model.Product{},
		model.User{},
		model.UserInfo{},
	)
}
