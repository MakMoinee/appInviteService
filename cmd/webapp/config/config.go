package config

import (
	"github.com/MakMoinee/appInviteService/internal/appInviteService/common"
	"github.com/spf13/viper"
)

var Registry *viper.Viper

func Set() {
	var err error
	viper.AddConfigPath(".")
	viper.AddConfigPath("../..")
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Registry = viper.GetViper()

	common.SERVER_PORT = Registry.GetString("SERVER_PORT")
	common.ENABLE_PROFILING = Registry.GetBool("ENABLE_PROFILING")
	common.DB_NAME = Registry.GetString("DB_NAME")
	common.DB_DRIVER = Registry.GetString("DB_DRIVER")
	common.MYSQL_USERNAME = Registry.GetString("MYSQL_USERNAME")
	common.MYSQL_PASSWORD = Registry.GetString("MYSQL_PASSWORD")
	common.CONNECTION_STRING = Registry.GetString("CONNECTION_STRING")
	common.RETRY_SLEEP = Registry.GetDuration("RETRY_SLEEP")
}
