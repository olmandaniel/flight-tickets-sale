package listener

import (
	"fmt"
	"log"

	domainEvents "github.com/olmandaniel/flight-tickets-sale/payment_service/event"
	"github.com/olmandaniel/flight-tickets-sale/payment_service/lib"
	"github.com/olmandaniel/flight-tickets-sale/payment_service/lib/msgqueue"
	"github.com/olmandaniel/flight-tickets-sale/payment_service/repository"
	"github.com/olmandaniel/flight-tickets-sale/payment_service/services"
)

type EventProcessor struct {
	msgqueue.EventListener
	*lib.Database
	repository.PaymentRepository
	msgqueue.EventEmitter
}

func (p *EventProcessor) ProcessEvents() error {
	log.Print("Listening to events ...")
	received, errors, err := p.EventListener.Listen("process.payment")
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
	case *domainEvents.ProcessPaymentEvent:
		usecase := services.NewProcessPayment(p.PaymentRepository, p.EventEmitter)
		usecase.Execute(*e)
	default:
		log.Printf("unknown event: %t", e)

	}
}
