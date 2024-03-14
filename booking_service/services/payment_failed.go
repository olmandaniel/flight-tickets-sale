package services

import (
	"github.com/olmandaniel/flight-tickets-sale/booking_service/event"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/lib/msgqueue"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/repository"
)

type PaymentFailed struct {
	repository.BookingRepository
	Emitter msgqueue.EventEmitter
}

func NewPaymentFailed(BookingRepository repository.BookingRepository, Emitter msgqueue.EventEmitter) *PaymentFailed {
	return &PaymentFailed{BookingRepository: BookingRepository, Emitter: Emitter}
}

func (usecase *PaymentFailed) Execute(bookId string) error {
	book, _ := usecase.BookingRepository.FindById(bookId)
	return usecase.Emitter.Emit(&event.MakeBookingFailureEvent{BookID: bookId, FlightID: book.FlightID, NumberOfSeats: book.NumberOfSeats})
}
