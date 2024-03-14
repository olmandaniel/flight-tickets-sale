package repository

import (
	"fmt"

	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/lib"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/model"
)

type FlightSearchRepository interface {
	SearchFlightByParams(origin string, destination string, date string) ([]model.Flight, error)
	GetFlightById(id string) (model.Flight, error)
	UpdateAvailableSeats(flightId string, numberOfSeats uint) error
}

type FlightSearchRepositoryImpl struct {
	DB *lib.Database
}

func NewFlightSearchRepository(DB *lib.Database) FlightSearchRepository {
	return &FlightSearchRepositoryImpl{DB: DB}
}

func (repo *FlightSearchRepositoryImpl) SearchFlightByParams(origin string, destination string, date string) ([]model.Flight, error) {
	var flights []model.Flight
	err := repo.DB.Model(&model.Flight{}).
		Joins("left join routes on routes.id = flights.route_id").Joins("inner join cities as origin on origin.id = routes.origin_id and origin.code = ?", origin).
		Joins("inner join cities as destination on destination.id = routes.destination_id and destination.code = ?", destination).Where("flights.date = ?", date).Find(&flights).Error
	return flights, err
}

func (repo *FlightSearchRepositoryImpl) GetFlightById(id string) (model.Flight, error) {
	flight := model.Flight{ID: id}
	result := repo.DB.Model(&model.Flight{}).Find(&flight)

	if result.Error != nil {
		return flight, result.Error
	}

	if result.RowsAffected == 0 {
		return flight, fmt.Errorf("the flight with id %s not found", id)
	}

	return flight, nil
}

func (repo *FlightSearchRepositoryImpl) UpdateAvailableSeats(flightId string, numberOfSeats uint) error {
	return repo.DB.Model(&model.Flight{}).Where("id = ?", flightId).Update("available_seats", numberOfSeats).Error
}
