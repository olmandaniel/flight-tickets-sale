package services

import (
	"fmt"

	domainEvents "github.com/olmandaniel/flight-tickets-sale/flight_search_service/event"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/lib/msgqueue"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/repository"
)

type ValidateBooking struct {
	repository.FlightSearchRepository
	Emitter msgqueue.EventEmitter
}

func NewValidateBooking(FlightSearchRepository repository.FlightSearchRepository, Emitter msgqueue.EventEmitter) *ValidateBooking {
	return &ValidateBooking{FlightSearchRepository: FlightSearchRepository, Emitter: Emitter}
}

func (usecase *ValidateBooking) Execute(event domainEvents.FlightBookedEvent) error {
	flight, err := usecase.FlightSearchRepository.GetFlightById(event.FlightID)
	if err != nil {
		usecase.Emitter.Emit(&domainEvents.BookingValidatedFailureEvent{})
		return err
	}

	if event.NumberOfSeats > flight.AvailableSeats {
		usecase.Emitter.Emit(&domainEvents.BookingValidatedFailureEvent{})
		return fmt.Errorf("the available seats excedeed")
	}

	availableSeatsAfterBooked := flight.AvailableSeats - event.NumberOfSeats
	err = usecase.FlightSearchRepository.UpdateAvailableSeats(event.FlightID, availableSeatsAfterBooked)
	if event.NumberOfSeats > flight.AvailableSeats {
		usecase.Emitter.Emit(&domainEvents.BookingValidatedFailureEvent{})
		return err
	}

	usecase.Emitter.Emit(&domainEvents.BookingValidatedSuccessfullyEvent{})
	return nil
}
