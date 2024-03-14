package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/handler"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/lib"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/lib/msgqueue/amqp"
	eventListener "github.com/olmandaniel/flight-tickets-sale/flight_search_service/listener"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/repository"
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

	flightRepository := repository.NewFlightSearchRepository(db)
	flightHandler := handler.NewFlightSearchHandler(flightRepository, responseService)
	r := app.Group("/flights")
	r.Get("/", flightHandler.SearchFlightByParams)
	r.Get("/:id", flightHandler.GetFlightById)

	processor := eventListener.EventProcessor{EventListener: listener, Database: db, FlightSearchRepository: flightRepository, EventEmitter: emitter}

	go processor.ProcessEvents()
	log.Fatal(app.Listen(fmt.Sprintf(":%s", env.AppPort)))
}
