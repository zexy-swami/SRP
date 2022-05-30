package parser

import "github.com/spf13/viper"

func ParseConfig(configName string) error {
	viper.AddConfigPath(".")
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}
