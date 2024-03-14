package listener

import (
	"fmt"
	"log"

	domainEvents "github.com/olmandaniel/flight-tickets-sale/booking_service/event"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/lib"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/lib/msgqueue"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/repository"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/services"
)

type EventProcessor struct {
	msgqueue.EventListener
	*lib.Database
	repository.BookingRepository
	Emitter msgqueue.EventEmitter
}

func (p *EventProcessor) ProcessEvents() error {
	log.Print("Listening to events ...")
	received, errors, err := p.EventListener.Listen("flight.booked.success", "flight.booked.failure", "process.payment.failed", "process.payment.success")
	if err != nil {
		fmt.Println(err)
		return err
	}

	for {
		select {
		case evt := <-received:
			fmt.Println(evt)
			p.handleEvent(evt)
		case err = <-errors:
			log.Printf("received error while processing msg: %s", err)
		}
	}
}

func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	case *domainEvents.BookingValidatedSuccessfullyEvent:
		log.Printf("event %s created: %v", e.BookID, e)
		makeBookingUsecase := services.NewMakeBooking(p.BookingRepository)
		makeBookingUsecase.Execute(e.BookID)
	case *domainEvents.BookingValidatedFailureEvent:
		log.Printf("event %s created: %v", e.BookID, e)
		cancelBookingUsecase := services.NewCancelBooking(p.BookingRepository)
		cancelBookingUsecase.Execute(e.BookID)
	case *domainEvents.ProcessPaymentSuccessfullyEvent:
		usecase := services.NewPaidBooking(p.BookingRepository)
		usecase.Execute(e.BookID)
	case *domainEvents.ProcessPaymentFailureEvent:
		usecase := services.NewPaymentFailed(p.BookingRepository, p.Emitter)
		usecase.Execute(e.BookID)
	default:
		log.Printf("unknown event: %t", e)

	}
}
