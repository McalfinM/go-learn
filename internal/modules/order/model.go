package order

import (
	"time"
)

type Booking struct {
	ID          int64  `json:"-"`
	BookingUUID string `json:"booking_uuid"`

	UserID   int64  `json:"-"`
	RoomUuid string `json:"-"`
	RoomID   int64  `json:"-"`

	BookingType string `json:"booking_type"`

	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`

	DurationHours *int `json:"duration_hours"`

	BasePrice     float64 `json:"base_price"`
	DepositAmount float64 `json:"deposit_amount"`

	Status string `json:"status"`

	GuestName string
	GuestKTP  string
}

type BookingAccess struct {
	ID         int64 `json:"-"`
	AccessUUID string

	BookingID int64
	RoomID    int64
	RoomUuid  string

	QRCode string

	ValidFrom  time.Time
	ValidUntil time.Time

	IsUsed bool
}
