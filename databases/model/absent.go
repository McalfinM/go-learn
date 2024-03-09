package model

import "time"

type Absent struct {
	ID        int
	Uuid      string
	User_uuid string
	CreatedAt time.Time
	UpdatedAt time.Time
}
