package constant

const (
	Exchange        = "booking.exchange"
	Queue           = "booking.queue"
	DeadLetterQueue = "booking.dlq"

	EventCreated   = "booking.created"
	EventCancelled = "booking.cancelled"
	EventExpired = "booking.expired"
)
