package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() {
	var err error
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")
	viper.SetConfigName("RefinedConfig")
	viper.SetConfigFile("biz/config/config.yml")
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", "8022")
	viper.SetDefault("database.sqlite.file", "biz/dal/database/refined.db")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error in reading config file: %s\n", err))
	}
}
