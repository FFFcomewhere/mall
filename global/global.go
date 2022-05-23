package global

import (
	"exam8/config"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Viper  *viper.Viper
)
