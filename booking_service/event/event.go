package event

type FlightBookedEvent struct {
	ID            string `json:"id"`
	FlightID      string `json:"flight_id"`
	NumberOfSeats uint   `json:"number_of_seats"`
}

func (*FlightBookedEvent) EventName() string {
	return "flight.booked"
}

type BookingValidatedSuccessfullyEvent struct {
	BookID string `json:"book_id"`
}

func (*BookingValidatedSuccessfullyEvent) EventName() string {
	return "flight.booked.success"
}

type BookingValidatedFailureEvent struct {
	BookID string `json:"book_id"`
}

func (*BookingValidatedFailureEvent) EventName() string {
	return "flight.booked.failure"
}

type MakeBookingFailureEvent struct {
	BookID        string `json:"book_id"`
	FlightID      string `json:"flight_id"`
	NumberOfSeats uint   `json:"number_of_seats"`
}

func (*MakeBookingFailureEvent) EventName() string {
	return "make.booking.failure"
}

type ProcessPaymentEvent struct {
	BookID string  `json:"book_id"`
	Price  float64 `json:"price"`
}

func (*ProcessPaymentEvent) EventName() string {
	return "make.booking.success"
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
