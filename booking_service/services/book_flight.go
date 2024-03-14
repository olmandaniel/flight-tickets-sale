package services

import (
	"github.com/olmandaniel/flight-tickets-sale/booking_service/dtos"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/event"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/lib/msgqueue"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/model"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/repository"
)

type BookFlight struct {
	repository.BookingRepository
	Emitter msgqueue.EventEmitter
}

func NewBookFlight(BookingRepository repository.BookingRepository, Emitter msgqueue.EventEmitter) *BookFlight {
	return &BookFlight{BookingRepository: BookingRepository, Emitter: Emitter}
}

func (usecase *BookFlight) Execute(flightBookRequest dtos.BookFlightRequest) error {
	totalPrice := float64(flightBookRequest.NumberOfSeats) * flightBookRequest.SeatPrice
	for _, v := range flightBookRequest.Luggages {
		totalPrice += v.Fear
	}

	passengers := make([]model.Passenger, 0)
	for _, v := range flightBookRequest.Passengers {
		passenger := model.Passenger{Name: v.Name, SeatNumber: *v.SeatNumber}
		passengers = append(passengers, passenger)
	}

	flightBooked := model.Booking{FlightID: flightBookRequest.FlightID, Status: "PENDING", TotalPrice: totalPrice, Passengers: passengers}
	bookId, err := usecase.BookingRepository.BookFlight(flightBooked)
	if err != nil {
		return err
	}

	err = usecase.Emitter.Emit(&event.FlightBookedEvent{ID: bookId, NumberOfSeats: flightBookRequest.NumberOfSeats})
	if err != nil {
		return err
	}

	return nil
}
