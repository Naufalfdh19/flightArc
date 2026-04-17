package bookingqueue

import (
	. "flight/modules/booking/constant"

	amqp "github.com/rabbitmq/amqp091-go"
)

func DeclareQueues(ch *amqp.Channel) error {
	if _, err := ch.QueueDeclare(DeadLetterQueue, true, false, false, false, nil); err != nil {
		return err
	}

	if err := ch.ExchangeDeclare(Exchange, "topic", true, false, false, false, nil); err != nil {
		return err
	}

	args := amqp.Table{
		"x-dead-letter-exchange":    "",
		"x-dead-letter-routing-key": DeadLetterQueue,
		"x-message-ttl": 90000,

	}
	
	if _, err := ch.QueueDeclare(Queue, true, false, false, false, args); err != nil {
		return err
	}

	for _, key := range []string{EventCreated, EventExpired} {
		if err := ch.QueueBind(Queue, key, Exchange, false, nil); err != nil {
			return err
		}
	}

	return nil
}