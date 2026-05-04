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

func (r *ProfileRepository) UpsertKTPImage(userID string, url string) error {
	_, err := r.DB.Exec(`
		INSERT INTO profiles (user_id, ktp_image_url)
		VALUES ($1, $2)
		ON CONFLICT (user_id)
		DO UPDATE SET ktp_image_url = EXCLUDED.ktp_image_url,
		              updated_at = NOW()
	`, userID, url)

	return err
}
