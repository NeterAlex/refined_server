package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func Init() {
	var err error
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	dir, err := os.Getwd()
	fmt.Println(dir)
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", "8022")
	viper.SetDefault("database.sql.file", "biz/dal/database/refined.db")
	viper.SetDefault("jwt.secretkey", "F54d|CX1v")

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error in reading config file: %s\n", err))
	}
}
