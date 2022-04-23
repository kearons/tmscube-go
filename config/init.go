package config

import (
	"github.com/spf13/viper"
)

var conf Config

type Config struct {
	Mysql Mysql `toml:mysql`

	Rabbitmq Rabbitmq `toml:rabbitmq`
}

type Mysql struct {
	Host     string
	Port     int
	Db       string
	Username string
	Password string
}

type Rabbitmq struct {
	Host     string
	Port     int
	Login    string
	Password string
}

func init() {
	viper.SetConfigName("dev")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}
}

func Load() {
}

func GetConf() Config {
	return conf
}
