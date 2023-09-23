package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DB DB `yaml:"db"`
}

type DB struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

var Conf = NewConfig()

func NewConfig() Config {
	env := os.Getenv("ENV")
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("app/config/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var congig Config
	err = viper.Unmarshal(&congig)
	if err != nil {
		panic(err)
	}

	return congig
}
