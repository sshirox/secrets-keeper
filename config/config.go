package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SafeWriteConfig()
}

func SaveToken(token string) {
	viper.Set("token", token)
	viper.WriteConfig()
}

func GetToken() string {
	return viper.GetString("token")
}
