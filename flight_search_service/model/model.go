package model

import "time"

type City struct {
	ID   string `gorm:"type:uuid;default:gen_random_uuid()"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Route struct {
	ID            string `gorm:"type:uuid;default:gen_random_uuid()"`
	OriginID      string `json:"origin_id"`
	Origin        City   `json:"origin"`
	DestinationID string `json:"destination_id"`
	Destination   City   `json:"destination"`
}

type Flight struct {
	ID             string     `gorm:"type:uuid;default:gen_random_uuid()"`
	Date           string     `json:"date"`
	DeapertureTime *time.Time `json:"deaperture_time"`
	ArrivalTime    *time.Time `json:"arrival_time"`
	Price          float64    `json:"price"`
	RouteID        string     `json:"route_id"`
	Route          Route      `json:"route"`
	AvailableSeats uint       `json:"available_seats"`
}

type FlightSeat struct {
	FlightID  string  `json:"flight_id"`
	SeatID    string  `json:"seat_id"`
	Fare      float64 `json:"fare"`
	Status    string  `json:"status"`
	BookingID string  `json:"booking_id"`
}

type Seat struct {
	ID     string `gorm:"type:uuid;default:gen_random_uuid()"`
	Number string `json:"number"`
	Type   string `json:"type"`
}
