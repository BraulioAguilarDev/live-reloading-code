package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Settings struct {
	AppName string
	Port    int
	Debug   bool
}

var Config Settings

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
	}
}
