package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

var Channel *amqp.Channel

func Connect(url string) error {
	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}

	Channel, err = conn.Channel()
	return err
}
