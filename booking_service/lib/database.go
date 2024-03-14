package lib

import (
	"fmt"
	"log"
	"time"

	"github.com/olmandaniel/flight-tickets-sale/booking_service/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(env Env) *Database {
	url := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable connect_timeout=5 timezone=Etc/Universal ",
		env.DBHost,
		env.DBPort,
		env.DBDatabase,
		env.DBUserName,
		env.DBPassword,
	)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Etc/Universal")
			return time.Now().In(ti)
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("set timezone=UTC")
	db.AutoMigrate(&model.Booking{}, &model.Passenger{})
	return &Database{db}
}
