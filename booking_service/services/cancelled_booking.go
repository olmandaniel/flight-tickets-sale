package services

import "github.com/olmandaniel/flight-tickets-sale/booking_service/repository"

type CancelledBooking struct {
	repository.BookingRepository
}

func NewCancelBooking(BookingRepository repository.BookingRepository) *CancelledBooking {
	return &CancelledBooking{BookingRepository: BookingRepository}
}

func (usecase *CancelledBooking) Execute(bookId string) error {
	return usecase.BookingRepository.CancelBooking(bookId)
}
