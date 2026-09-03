package config

import "github.com/spf13/viper"

func InitViper() {
	viper.SetConfigFile(".env")

	_ = viper.ReadInConfig()

	viper.AutomaticEnv()
}
