package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.AddConfigPath(".")
	viper.AddConfigPath("config/")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Config not found, creating new one...")
		viper.SafeWriteConfig()
	}
}

func SaveToken(token string) {
	viper.Set("token", token)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println("Error writing config:", err)
	}
}

func GetToken() string {
	token := viper.GetString("token")
	log.Println("Loaded token:", token)
	return token
}
