package event

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

type FlightBookedEvent struct {
	ID            string `json:"id"`
	FlightID      string `json:"flight_id"`
	NumberOfSeats uint   `json:"number_of_seats"`
}

func (*FlightBookedEvent) EventName() string {
	return "flight.booked"
}

type MakeBookingFailureEvent struct {
	BookID        string `json:"book_id"`
	FlightID      string `json:"flight_id"`
	NumberOfSeats uint   `json:"number_of_seats"`
}

func (*MakeBookingFailureEvent) EventName() string {
	return "make.booking.failure"
}
