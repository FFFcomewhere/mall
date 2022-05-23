package initialize

import (
	"exam8/global"
	"fmt"
	"github.com/jinzhu/gorm"
)

func Mysql() {
	var err error
	mysqlCfg := global.Config.Mysql

	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlCfg.User,
		mysqlCfg.Password,
		mysqlCfg.Host,
		mysqlCfg.Name,
	)

	global.DB, err = gorm.Open(mysqlCfg.Type, args)
	if err != nil {
		fmt.Println("数据库连接失败 err: ", err)
	} else {
		fmt.Println("数据库连接成功")
	}

	//设置数据库的线程池
	global.DB.DB().SetMaxIdleConns(10)  //设置空闲连接池中的最大连接数
	global.DB.DB().SetMaxOpenConns(100) //设置与数据库的最大打开连接数

	//是否显示SQL语句
	global.DB.LogMode(true)

}
