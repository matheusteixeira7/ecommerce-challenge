package config

import (
	"github.com/spf13/viper"
)

type Conf struct {
	PORT        string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_SSL_MODE string
}

func LoadConfig() (*Conf, error) {
	var cfg *Conf

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath("..")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
