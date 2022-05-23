package main

import (
	"exam8/global"
	"exam8/initialize"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	initialize.Config()
	initialize.Mysql()
	initialize.DBTables()
}

func main() {
	defer global.DB.Close()

	////使用标准库开启服务,未完成
	//serverCfg := global.Config.Server
	//gin.SetMode(serverCfg.RunMode)
	//server := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", serverCfg.HttpPort),
	//	Handler:        initialize.Router(),
	//	ReadTimeout:    serverCfg.ReadTimeout,
	//	WriteTimeout:   serverCfg.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//server.ListenAndServe()

	//使用gin开启服务
	r := initialize.Router()
	//global.DB = logger.Default.LogMode(logger.Silent)
	r.Run(":8080")
}
