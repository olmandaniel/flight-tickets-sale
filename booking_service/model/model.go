package model

type Booking struct {
	ID            string      `gorm:"type:uuid;default:gen_random_uuid()"`
	FlightID      string      `json:"flight_id"`
	Status        string      `json:"status"`
	TotalPrice    float64     `json:"total_price"`
	Passengers    []Passenger `json:"passengers"`
	NumberOfSeats uint        `json:"number_of_seats"`
}

type Passenger struct {
	ID         string  `gorm:"type:uuid;default:gen_random_uuid()"`
	BookingID  string  `json:"booking_id"`
	Booking    Booking `json:"booking"`
	Name       string  `json:"name"`
	SeatNumber string  `json:"seat_number"`
}
