package order

import (
	"database/sql"
	"time"
)

type BookingRepository struct {
	DB *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{DB: db}
}

func (r *BookingRepository) Create(tx *sql.Tx, b *Booking) error {
	query := `
	INSERT INTO bookings (
		user_id,
		room_id,
		booking_type,
		start_time,
		end_time,
		duration_hours,
		base_price,
		deposit_amount,
		status
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	RETURNING id, booking_uuid
	`

	return tx.QueryRow(
		query,
		b.UserID,
		b.RoomID,
		b.BookingType,
		b.StartTime,
		b.EndTime,
		b.DurationHours,
		b.BasePrice,
		b.DepositAmount,
		b.Status,
	).Scan(&b.ID, &b.BookingUUID)
}

func (r *BookingRepository) CheckOverlap(
	tx *sql.Tx,
	roomID int64,
	start time.Time,
	end time.Time,
) (bool, error) {

	query := `
	SELECT 1
	FROM bookings
	WHERE room_id = $1
	AND status IN ('paid','active')
	AND (
		(start_time, end_time) OVERLAPS ($2, $3)
	)
	LIMIT 1
	`

	var exist int
	err := tx.QueryRow(query, roomID, start, end).Scan(&exist)

	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *BookingRepository) CreateAccess(tx *sql.Tx, a *BookingAccess) error {
	query := `
	INSERT INTO booking_accesses (
		booking_id,
		room_id,
		qr_code,
		valid_from,
		valid_until
	)
	VALUES ($1,$2,$3,$4,$5)
	RETURNING id, access_uuid
	`

	return tx.QueryRow(
		query,
		a.BookingID,
		a.RoomID,
		a.QRCode,
		a.ValidFrom,
		a.ValidUntil,
	).Scan(&a.ID, &a.AccessUUID)
}

func (r *BookingRepository) FindAccessByQR(qr string) (*BookingAccess, error) {
	query := `
	SELECT 
		id,
		booking_id,
		room_id,
		qr_code,
		valid_from,
		valid_until,
		is_used
	FROM booking_accesses
	WHERE qr_code=$1
	`

	var a BookingAccess

	err := r.DB.QueryRow(query, qr).Scan(
		&a.ID,
		&a.BookingID,
		&a.RoomID,
		&a.QRCode,
		&a.ValidFrom,
		&a.ValidUntil,
		&a.IsUsed,
	)

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *BookingRepository) UpdateStatus(tx *sql.Tx, bookingID int64, status string) error {
	_, err := tx.Exec(`
		UPDATE bookings
		SET status=$1
		WHERE id=$2
	`, status, bookingID)

	return err
}

func (r *BookingRepository) FindById(room_id string) (*BookingAccess, error) {
	query := `
	SELECT 
		id,
		booking_id,
		room_id,
		qr_code,
		valid_from,
		valid_until,
		is_used
	FROM booking_accesses
	WHERE room_id=$1
	`

	var a BookingAccess

	err := r.DB.QueryRow(query, room_id).Scan(
		&a.ID,
		&a.BookingID,
		&a.RoomID,
		&a.QRCode,
		&a.ValidFrom,
		&a.ValidUntil,
		&a.IsUsed,
	)

	if err != nil {
		return nil, err
	}

	return &a, nil
}
