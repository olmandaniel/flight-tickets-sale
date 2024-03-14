package lib

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppPort        string `mapstructure:"APP_PORT"`
	DBHost         string `mapstructure:"DATABASE_HOST"`
	DBUserName     string `mapstructure:"DATABASE_USER_NAME"`
	DBDatabase     string `mapstructure:"DATABASE_NAME"`
	DBPort         string `mapstructure:"DATABASE_PORT"`
	DBPassword     string `mapstructure:"DATABASE_PASSWORD"`
	BrokerDriver   string `mapstructure:"BROKER_DRIVER"`
	BrokerHost     string `mapstructure:"BROKER_HOST"`
	BrokerPort     string `mapstructure:"BROKER_PORT"`
	BrokerUsername string `mapstructure:"BROKER_USERNAME"`
	BrokerPassword string `mapstructure:"BROKER_PASSWORD"`
}

func NewEnv() Env {
	env := Env{}

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Cannot read config file : ", err)
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded :", err)
	}

	return env
}
