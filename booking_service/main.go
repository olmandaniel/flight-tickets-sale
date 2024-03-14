package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/handler"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/lib"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/lib/msgqueue/amqp"
	eventListener "github.com/olmandaniel/flight-tickets-sale/booking_service/listener"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/repository"
)

func main() {
	app := fiber.New()

	env := lib.NewEnv()
	db := lib.NewDatabase(env)
	responseService := lib.NewResponseService()
	messageBroker := lib.NewBroker(env)
	emitter, err := amqp.NewAMQPEventEmitter(messageBroker.Connection)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := amqp.NewAMQPEventListener(messageBroker.Connection, "orders")
	if err != nil {
		log.Fatal(err)
	}
	bookingRepository := repository.NewBookingRepository(db)

	processor := eventListener.EventProcessor{EventListener: listener, Database: db, BookingRepository: bookingRepository, Emitter: emitter}

	bookingHandler := handler.NewBookingHandler(emitter, responseService, bookingRepository)

	r := app.Group("/book-flights")
	r.Post("/", bookingHandler.BookFlight)

	go processor.ProcessEvents()
	log.Fatal(app.Listen(fmt.Sprintf(":%s", env.AppPort)))
}
