package services

import "github.com/olmandaniel/flight-tickets-sale/booking_service/repository"

type PaidBooking struct {
	repository.BookingRepository
}

func NewPaidBooking(BookingRepository repository.BookingRepository) *PaidBooking {
	return &PaidBooking{BookingRepository: BookingRepository}
}

func (usecase *PaidBooking) Execute(bookId string) error {
	return usecase.BookingRepository.PaidBooking(bookId)
}
