package account

import (
	"database/sql"
)

type ProfileRepository struct {
	DB *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{DB: db}
}

func (r *ProfileRepository) FindByUserID(userID int64) (*Profile, error) {
	var p Profile

	err := r.DB.QueryRow(`
		SELECT id, profile_uuid, user_id, full_name, ktp_number, phone, address, date_of_birth
		FROM profiles
		WHERE user_id=$1
	`, userID).Scan(
		&p.ID,
		&p.ProfileUUID,
		&p.UserID,
		&p.FullName,
		&p.KTP,
		&p.Phone,
		&p.Address,
		&p.DateOfBirth,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}
