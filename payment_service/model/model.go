package model

import "time"

type Payment struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid()"`
	BookingID string    `json:"booking_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	Date      time.Time `json:"payment_date"`
}

func (*Payment) TableName() string {
	return "payments"
}
