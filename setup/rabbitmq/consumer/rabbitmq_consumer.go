package consumer

import (
	"flight/setup/router"

	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func SetupRabbitMQConsumer(db *gorm.DB, mqch *amqp.Channel) {
	repos := router.SetupRepo(db)
	services := router.SetupService(repos, mqch)
	go func() {
		services.BookingService.CheckExpiredBooking()
	}()
}
