package config

import (
	"github.com/spf13/viper"
	"log"
)

var cfgDatabase *viper.Viper
var cfgApplication *viper.Viper
var cfgMessage *viper.Viper


func init() {
	viper.SetConfigName("settings")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	cfgDatabase = viper.Sub("settings.database")
	if cfgDatabase == nil {
		panic("config not found settings.database")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)


	cfgApplication = viper.Sub("settings.application")
	if cfgApplication == nil {
		panic("config not found settings.application")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	cfgMessage = viper.Sub("settings.message")
	if cfgMessage == nil {
		panic("config not found settings.message")
	}
	MessageConfig = InitMessage(cfgMessage)
}

func SetApplicationIsInit() {
	SetConfig("./config","settings.application.isInit", false)
}

func SetConfig(configPath string,key string,value interface{}){
	viper.AddConfigPath(configPath)
	viper.Set(key, value)
	viper.WriteConfig()
}
