package services

import (
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/model"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/repository"
)

type searchFlightByParams struct {
	repository.FlightSearchRepository
}

func NewSearchFlightByParams(FlightSearchRepository repository.FlightSearchRepository) *searchFlightByParams {
	return &searchFlightByParams{FlightSearchRepository: FlightSearchRepository}
}

func (u *searchFlightByParams) Execute(origin string, destination string, date string) ([]model.Flight, error) {
	return u.FlightSearchRepository.SearchFlightByParams(origin, destination, date)
}
