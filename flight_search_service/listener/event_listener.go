package listener

import (
	"fmt"
	"log"

	domainEvents "github.com/olmandaniel/flight-tickets-sale/flight_search_service/event"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/lib"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/lib/msgqueue"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/repository"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/services"
)

type EventProcessor struct {
	msgqueue.EventListener
	*lib.Database
	repository.FlightSearchRepository
	msgqueue.EventEmitter
}

func (p *EventProcessor) ProcessEvents() error {
	log.Print("Listening to events ...")
	received, errors, err := p.EventListener.Listen("flight.booked", "make.booking.failure")
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
	case *domainEvents.FlightBookedEvent:
		usecase := services.NewValidateBooking(p.FlightSearchRepository, p.EventEmitter)
		usecase.Execute(*e)
	case *domainEvents.MakeBookingFailureEvent:
		usecase := services.NewRevertBookingSeats(p.FlightSearchRepository, p.EventEmitter)
		usecase.Execute(*e)
	default:
		log.Printf("unknown event: %t", e)

	}
}
