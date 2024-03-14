package event

type ProcessPaymentEvent struct {
	BookID string  `json:"book_id"`
	Price  float64 `json:"price"`
}

func (*ProcessPaymentEvent) EventName() string {
	return "process.payment"
}

type ProcessPaymentFailureEvent struct {
	BookID string `json:"book_id"`
}

func (*ProcessPaymentFailureEvent) EventName() string {
	return "process.payment.failed"
}

type ProcessPaymentSuccessfullyEvent struct {
	BookID string `json:"book_id"`
}

func (*ProcessPaymentSuccessfullyEvent) EventName() string {
	return "process.payment.success"
}
