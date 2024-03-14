package lib

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppPort             string `mapstructure:"APP_PORT"`
	SearchFlightService string `mapstructure:"API_FLIGHT_SERVICE"`
	BookingService      string `mapstructure:"API_BOOKING_SERVICE"`
	PaymentService      string `mapstructure:"API_PAYMENT_SERVICE"`
}

func NewEnv() Env {
	env := Env{}

	viper.SetConfigFile("./.env")
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
