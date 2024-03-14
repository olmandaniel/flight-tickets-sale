package repository

import (
	"github.com/olmandaniel/flight-tickets-sale/payment_service/lib"
	"github.com/olmandaniel/flight-tickets-sale/payment_service/model"
)

type PaymentRepository interface {
	CreatePayment(data model.Payment) error
}

type PaymentRepositoryImpl struct {
	DB lib.Database
}

func NewPaymentRepository(DB lib.Database) PaymentRepository {
	return &PaymentRepositoryImpl{DB: DB}
}

func (repo *PaymentRepositoryImpl) CreatePayment(data model.Payment) error {
	return repo.DB.Model(&model.Payment{}).Create(&data).Error
}
