package initialize

import (
	"exam8/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

const defaultConfigFile = "config/config.yaml"

func Config() {
	global.Viper = viper.New()
	global.Viper.SetConfigFile(defaultConfigFile)
	err := global.Viper.ReadInConfig()
	if err != nil {
		fmt.Println(" Fatal error config file: ", err)
	}

	global.Viper.WatchConfig()

	//将json格式的配置文件转成字符串
	global.Viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)

		if err := global.Viper.Unmarshal(&global.Config); err != nil {
			fmt.Println(" err: ", err)
		}
	})

	if err := global.Viper.Unmarshal(&global.Config); err != nil {
		fmt.Println(" err: ", err)
	}

	//初始化配置
	serverCfg := &global.Config.Server
	serverCfg.ReadTimeout = serverCfg.ReadTimeout * time.Second
	serverCfg.WriteTimeout = serverCfg.WriteTimeout * time.Second

}
