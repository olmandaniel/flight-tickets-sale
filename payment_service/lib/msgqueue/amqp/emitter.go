package amqp

import (
	"context"
	"encoding/json"

	"github.com/olmandaniel/flight-tickets-sale/payment_service/lib/msgqueue"
	amqp "github.com/rabbitmq/amqp091-go"
)

type amqpEventEmitter struct {
	*amqp.Connection
}

func (e *amqpEventEmitter) setup() error {
	channel, err := e.Connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()
	channel.QueueDeclare("orders", true, false, false, false, amqp.Table{})
	return channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)
}

func NewAMQPEventEmitter(conn *amqp.Connection) (msgqueue.EventEmitter, error) {
	emitter := &amqpEventEmitter{Connection: conn}

	err := emitter.setup()
	if err != nil {
		return emitter, err
	}

	return emitter, nil
}

func (e *amqpEventEmitter) Emit(event msgqueue.Event) error {
	jsonDoc, err := json.Marshal(event)
	if err != nil {
		return err
	}
	channel, err := e.Connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	msg := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": event.EventName()},
		Body:        jsonDoc,
		ContentType: "application/json",
	}

	return channel.PublishWithContext(context.Background(), "events", event.EventName(), false, false, msg)
}
