package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/olmandaniel/flight-tickets-sale/payment_service/lib"
	"github.com/olmandaniel/flight-tickets-sale/payment_service/lib/msgqueue/amqp"
	eventListener "github.com/olmandaniel/flight-tickets-sale/payment_service/listener"
	"github.com/olmandaniel/flight-tickets-sale/payment_service/repository"
)

func main() {
	app := fiber.New()

	env := lib.NewEnv()
	db := lib.NewDatabase(env)
	messageBroker := lib.NewBroker(env)
	emitter, err := amqp.NewAMQPEventEmitter(messageBroker.Connection)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := amqp.NewAMQPEventListener(messageBroker.Connection, "orders")
	if err != nil {
		log.Fatal(err)
	}
	paymentRepository := repository.NewPaymentRepository(db)
	processor := eventListener.EventProcessor{EventListener: listener, Database: &db, PaymentRepository: paymentRepository, EventEmitter: emitter}

	go processor.ProcessEvents()
	log.Fatal(app.Listen(fmt.Sprintf(":%s", env.AppPort)))
}
