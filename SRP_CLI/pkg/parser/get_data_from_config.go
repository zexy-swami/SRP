package parser

import "github.com/spf13/viper"

func GetDataFromConfig(key string) string {
	return viper.GetString(key)
}
