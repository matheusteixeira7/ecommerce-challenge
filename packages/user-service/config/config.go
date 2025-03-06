package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Conf struct {
	PORT string
}

func LoadConfig() (*Conf, error) {
	var cfg *Conf
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter diretório root: %v", err)
	}

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(rootDir)
	viper.SetConfigFile(rootDir + "/.env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
