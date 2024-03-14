package services

import (
	"time"

	domainEvents "github.com/olmandaniel/flight-tickets-sale/payment_service/event"
	"github.com/olmandaniel/flight-tickets-sale/payment_service/lib/msgqueue"
	"github.com/olmandaniel/flight-tickets-sale/payment_service/model"
	"github.com/olmandaniel/flight-tickets-sale/payment_service/repository"
)

type ProcessPayment struct {
	repository.PaymentRepository
	Emitter msgqueue.EventEmitter
}

func NewProcessPayment(PaymentRepository repository.PaymentRepository, Emitter msgqueue.EventEmitter) *ProcessPayment {
	return &ProcessPayment{PaymentRepository: PaymentRepository, Emitter: Emitter}
}

func (usecase *ProcessPayment) Execute(event domainEvents.ProcessPaymentEvent) error {
	payment := model.Payment{BookingID: event.BookID, Amount: event.Price, Status: "PENDING", Date: time.Now()}
	err := usecase.PaymentRepository.CreatePayment(payment)
	if err != nil {
		usecase.Emitter.Emit(&domainEvents.ProcessPaymentFailureEvent{BookID: event.BookID})
		return err
	}

	factory := PaymentProcessorFactory{}
	proccessor := factory.GetPaymentDriver("CREDIT_CARD")
	err = proccessor.Pay()
	if err != nil {
		usecase.Emitter.Emit(&domainEvents.ProcessPaymentFailureEvent{BookID: event.BookID})
		return err
	}

	usecase.Emitter.Emit(&domainEvents.ProcessPaymentSuccessfullyEvent{BookID: event.BookID})
	return nil
}
