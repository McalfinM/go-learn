package account

import (
	"time"
)

type Role struct {
	RoleID   int    `json:"id"`
	RoleUuid string `json:"role_uuid"`
	Name     string `json:"name"`
}

type User struct {
	ID       int64  `json:"-"`
	UserUUID string `json:"user_uuid"`

	Email    string `json:"email"`
	Password string `json:"-"`

	RoleID   int64 `json:"-"`
	RoleName string
}

type Profile struct {
	ID          int64  `json:"-"`
	ProfileUUID string `json:"profile_uuid"`

	UserID int64 `json:"-"`

	FullName string `json:"full_name"`
	KTP      string `json:"ktp_number"`
	Phone    string `json:"phone"`

	Address     string     `json:"address"`
	DateOfBirth *time.Time `json:"date_of_birth"`

	KtpImageUrl string `json:"-"`
}
