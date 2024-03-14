package services

import (
	domainEvents "github.com/olmandaniel/flight-tickets-sale/flight_search_service/event"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/lib/msgqueue"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/repository"
)

type RevertBookingSeats struct {
	repository.FlightSearchRepository
	Emitter msgqueue.EventEmitter
}

func NewRevertBookingSeats(FlightSearchRepository repository.FlightSearchRepository, Emitter msgqueue.EventEmitter) *RevertBookingSeats {
	return &RevertBookingSeats{FlightSearchRepository: FlightSearchRepository, Emitter: Emitter}
}

func (usecase *RevertBookingSeats) Execute(event domainEvents.MakeBookingFailureEvent) error {
	flight, err := usecase.FlightSearchRepository.GetFlightById(event.FlightID)
	if err != nil {
		usecase.Emitter.Emit(&domainEvents.BookingValidatedFailureEvent{BookID: event.BookID})
		return err
	}

	flight.AvailableSeats = flight.AvailableSeats + event.NumberOfSeats
	err = usecase.FlightSearchRepository.UpdateAvailableSeats(event.FlightID, flight.AvailableSeats)
	usecase.Emitter.Emit(&domainEvents.BookingValidatedFailureEvent{BookID: event.BookID})
	return err
}
