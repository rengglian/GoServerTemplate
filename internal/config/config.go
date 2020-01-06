package config

import (
	"GoServerTemplate/internal/check"
	"GoServerTemplate/internal/structs"

	"github.com/spf13/viper"
)

//Init of gamemaster
func Init() structs.Configuration {
	serverConf := viper.New()
	serverConf.SetConfigName("config")    // name of config file (without extension)
	serverConf.AddConfigPath("./config/") // optionally look for config in the working directory
	err := serverConf.ReadInConfig()      // Find and read the config file

	check.Error("ReadConfig failed", err)

	var config structs.Configuration
	err = serverConf.Unmarshal(&config)
	check.Error("Unmarshal failed", err)

	return config
}
