package constant

const (
	Exchange        = "bookings.exchange"
	Queue           = "bookings.queue"
	DeadLetterQueue = "bookings.dlq"

	EventCreated   = "bookings.created"
	EventCancelled = "bookings.cancelled"
	EventExpired = "bookings.expired"
)
