package service

import (
	"encoding/json"
	"flight/modules/booking/constant"
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

	// Publish ke antrian PENDING (Bukan yang timeout_queue!)
	// Karena antrian ini yang punya x-message-ttl 15 menit
	err = s.ch.Publish(
		"",                  // exchange (kosongkan jika pakai default direct)
		constant.Queue, // routing key (nama antrian waiting room kamu)
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Printf("Gagal mengirim pesan timeout ke RabbitMQ: %v", err)
		return err
	}

	log.Printf("Pesan timeout untuk booking %s berhasil dikirim", bookingID)
	return nil
}