package lib

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Broker struct {
	*amqp.Connection
}

func NewBroker(env Env) *Broker {
	url := fmt.Sprintf("%s://%s:%s@%s:%s", env.BrokerDriver, env.BrokerUsername, env.BrokerPassword, env.BrokerHost, env.BrokerPort)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal("Coud not establish AMQP connection: " + err.Error())
	}

	// defer conn.Close()
	return &Broker{Connection: conn}
}
