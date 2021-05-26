package config

import (
	"log"

	"github.com/spf13/viper"
)

func GetVar(key string) string {

	viper.SetConfigFile("/var/www/api.govirtuoso.org/bin/.env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error loading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value

}
