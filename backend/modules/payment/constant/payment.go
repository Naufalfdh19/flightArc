package constant

const (
	CREATE_PAYMENT          = "create payment"
	GET_PAYMENT_BY_ID       = "get payment by id"
	HANDLE_PAYMENT_CALLBACK = "handle payment callback"
)

const (
	ProviderManualVA = "MANUAL_VA"
	MethodVirtualVA  = "VIRTUAL_ACCOUNT"
)

const (
	StatusPending = "PENDING"
	StatusPaid    = "PAID"
	StatusFailed  = "FAILED"
	StatusExpired = "EXPIRED"
)
