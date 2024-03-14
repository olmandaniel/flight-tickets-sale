package repository

import (
	"fmt"

	"github.com/olmandaniel/flight-tickets-sale/booking_service/lib"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/model"
)

type BookingRepository interface {
	BookFlight(data model.Booking) (string, error)
	CancelBooking(bookId string) error
	MakeBooking(bookId string) error
	FindById(bookId string) (model.Booking, error)
	PaidBooking(bookId string) error
}

type BookingRepositoryImpl struct {
	DB *lib.Database
}

func NewBookingRepository(DB *lib.Database) BookingRepository {
	return &BookingRepositoryImpl{DB: DB}
}

func (repo *BookingRepositoryImpl) BookFlight(data model.Booking) (string, error) {
	result := repo.DB.Create(&data)
	return data.ID, result.Error
}

func (repo *BookingRepositoryImpl) CancelBooking(bookId string) error {
	return repo.DB.Model(&model.Booking{}).Where("id = ?", bookId).Update("status = ?", "CANCELLED").Error
}

func (repo *BookingRepositoryImpl) MakeBooking(bookId string) error {
	return repo.DB.Model(&model.Booking{}).Where("id = ?", bookId).Update("status = ?", "BOOKING").Error
}

func (repo *BookingRepositoryImpl) PaidBooking(bookId string) error {
	return repo.DB.Model(&model.Booking{}).Where("id = ?", bookId).Update("status = ?", "COMPLETED").Error
}

func (repo *BookingRepositoryImpl) FindById(bookId string) (model.Booking, error) {
	booking := model.Booking{ID: bookId}
	result := repo.DB.Model(&model.Booking{}).Find(&booking)

	if result.RowsAffected == 0 {
		return booking, fmt.Errorf("booking with id %s not found", bookId)
	}

	if result.Error != nil {
		return booking, result.Error
	}

	return booking, nil
}
