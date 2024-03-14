package services

import (
	"github.com/olmandaniel/flight-tickets-sale/booking_service/event"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/lib/msgqueue"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/repository"
)

type MakeBooking struct {
	repository.BookingRepository
	Emitter msgqueue.EventEmitter
}

func NewMakeBooking(BookingRepository repository.BookingRepository) *MakeBooking {
	return &MakeBooking{BookingRepository: BookingRepository}
}

func (usecase *MakeBooking) Execute(bookId string) error {
	booking, err := usecase.BookingRepository.FindById(bookId)
	if err != nil {
		usecase.Emitter.Emit(&event.MakeBookingFailureEvent{FlightID: booking.FlightID, NumberOfSeats: booking.NumberOfSeats, BookID: bookId})
		return err
	}

	err = usecase.BookingRepository.MakeBooking(bookId)
	if err != nil {
		usecase.Emitter.Emit(&event.MakeBookingFailureEvent{FlightID: booking.FlightID, NumberOfSeats: booking.NumberOfSeats, BookID: bookId})
		return err
	}

	usecase.Emitter.Emit(&event.ProcessPaymentEvent{BookID: bookId, Price: booking.TotalPrice})
	return nil
}
