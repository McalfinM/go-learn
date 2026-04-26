package room

import "time"

type Room struct {
	RoomID   string  `json:"room_id"`
	Name     string  `json:"name"`
	Location *string `json:"location,omitempty"`

	PriceTransit float64 `json:"price_transit"`
	PriceDaily   float64 `json:"price_daily"`
	PriceMonthly float64 `json:"price_monthly"`

	DepositAmount *float64 `json:"deposit_amount,omitempty"`

	IsActive *bool `json:"is_active,omitempty"`

	IsTransitAvailable *bool `json:"is_transit_available,omitempty"`
	IsDailyAvailable   *bool `json:"is_daily_available,omitempty"`
	IsMonthlyAvailable *bool `json:"is_monthly_available,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}
