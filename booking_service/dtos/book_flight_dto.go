package dtos

type Passenger struct {
	Name       string  `json:"name"`
	SeatNumber *string `json:"seat_number"`
}

type Luggage struct {
	Type string  `json:"type"`
	Fear float64 `json:"fear"`
}

type BookFlightRequest struct {
	FlightID      string      `json:"flight_id"`
	SeatPrice     float64     `json:"seat_price"`
	NumberOfSeats uint        `json:"number_of_seats"`
	Passengers    []Passenger `json:"passengers"`
	Luggages      []Luggage   `json:"luggages"`
}
