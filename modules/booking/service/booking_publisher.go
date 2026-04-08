package service

import (
	"encoding/json"
	. "flight/modules/booking/constant"
	"log"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (s *BookingServiceImpl) PublishBookingTimeout(bookingID uuid.UUID) error {
	// Konversi UUID ke JSON byte
	body, err := json.Marshal(bookingID)
	if err != nil {
		return err
	}

	err = s.ch.Publish(
		Exchange,                  // exchange (kosongkan jika pakai default direct)
		Queue, // routing key (nama antrian waiting room kamu)
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Printf("Failed to send a message to RabbitMQ: %v", err)
		return err
	}

	log.Printf("Timeout message %s successfully sent", bookingID)
	return nil
}