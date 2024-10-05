package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Config = viper.New()

func init() {
	Config.AddConfigPath("conf")
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.WatchConfig()
	err := Config.ReadInConfig()
	if err != nil {
		zap.L().Fatal("Read config failed", zap.Error(err))
	}
}
