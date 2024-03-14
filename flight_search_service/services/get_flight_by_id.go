package services

import (
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/model"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/repository"
)

type getFlightById struct {
	repository.FlightSearchRepository
}

func NewGetFlightById(FlightSearchRepository repository.FlightSearchRepository) *getFlightById {
	return &getFlightById{FlightSearchRepository: FlightSearchRepository}
}

func (u *getFlightById) Execute(id string) (model.Flight, error) {
	return u.FlightSearchRepository.GetFlightById(id)
}
