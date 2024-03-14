package amqp

import (
	"encoding/json"
	"fmt"

	domainEvents "github.com/olmandaniel/flight-tickets-sale/payment_service/event"
	"github.com/olmandaniel/flight-tickets-sale/payment_service/lib/msgqueue"
	amqp "github.com/rabbitmq/amqp091-go"
)

type amqpEventListener struct {
	*amqp.Connection
	queue string
}

func (l *amqpEventListener) setup() error {
	channel, err := l.Connection.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	_, err = channel.QueueDeclare(l.queue, true, false, false, false, nil)
	return err
}

func NewAMQPEventListener(conn *amqp.Connection, queue string) (msgqueue.EventListener, error) {
	listener := &amqpEventListener{Connection: conn, queue: queue}

	err := listener.setup()
	if err != nil {
		return nil, err
	}

	return listener, nil
}

func (l *amqpEventListener) Listen(eventNames ...string) (<-chan msgqueue.Event, <-chan error, error) {
	channel, err := l.Connection.Channel()
	if err != nil {
		return nil, nil, err
	}

	// defer channel.Close()

	for _, eventName := range eventNames {
		if err := channel.QueueBind(l.queue, eventName, "events", false, nil); err != nil {
			return nil, nil, err
		}
	}

	msgs, err := channel.Consume(l.queue, "", false, false, false, false, nil)
	if err != nil {
		return nil, nil, err
	}

	events := make(chan msgqueue.Event)
	errors := make(chan error)

	go func() {
		for msg := range msgs {
			rawEventName, ok := msg.Headers["x-event-name"]
			if !ok {
				errors <- fmt.Errorf("msg did'nt contain x-event-name header")
				msg.Nack(false, false)
				continue
			}

			eventName, ok := rawEventName.(string)
			if !ok {
				errors <- fmt.Errorf("x-event-name header is not string, but %t", rawEventName)
				msg.Nack(false, false)
				continue
			}

			var event msgqueue.Event
			switch eventName {
			case "process.payment":
				event = new(domainEvents.ProcessPaymentEvent)
			default:
				errors <- fmt.Errorf("event type %s is unknown", eventName)
				continue
			}

			err := json.Unmarshal(msg.Body, event)
			if err != nil {
				errors <- err
				continue
			}

			events <- event

		}
	}()

	return events, errors, nil
}
