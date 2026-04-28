package room

import (
	"database/sql"
	"go-api/internal/utils"
	"strconv"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Create(room *Room) error {
	query := `
	INSERT INTO rooms (
		name, location,
		price_transit, price_daily, price_monthly,
		deposit_amount,
		is_active,
		is_transit_available, is_daily_available, is_monthly_available
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	RETURNING room_id, created_at
	`

	return r.DB.QueryRow(
		query,
		room.Name,
		room.Location,
		room.PriceTransit,
		room.PriceDaily,
		room.PriceMonthly,
		room.DepositAmount,
		room.IsActive,
		room.IsTransitAvailable,
		room.IsDailyAvailable,
		room.IsMonthlyAvailable,
	).Scan(&room.RoomID, &room.CreatedAt)
}

func (r *Repository) FindAll(filters map[string]interface{}) ([]Room, error) {
	query := `
	SELECT 
		room_id,
		name,
		location,
		price_transit,
		price_daily,
		price_monthly,
		deposit_amount,
		is_active,
		is_transit_available,
		is_daily_available,
		is_monthly_available,
		created_at
	FROM rooms
	WHERE 1=1
	`

	args := []interface{}{}
	i := 1

	// filter name (LIKE)
	if v, ok := filters["name"]; ok {
		query += " AND name ILIKE $" + strconv.Itoa(i)
		args = append(args, "%"+v.(string)+"%")
		i++
	}

	// filter is_active
	if v, ok := filters["is_active"]; ok {
		query += " AND is_active=$" + strconv.Itoa(i)
		args = append(args, v)
		i++
	}

	// price range
	if v, ok := filters["price_daily_min"]; ok {
		query += " AND price_daily >= $" + strconv.Itoa(i)
		args = append(args, v)
		i++
	}

	if v, ok := filters["price_daily_max"]; ok {
		query += " AND price_daily <= $" + strconv.Itoa(i)
		args = append(args, v)
		i++
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []Room

	for rows.Next() {
		var room Room

		err := rows.Scan(
			&room.RoomID,
			&room.Name,
			&room.Location,
			&room.PriceTransit,
			&room.PriceDaily,
			&room.PriceMonthly,
			&room.DepositAmount,
			&room.IsActive,
			&room.IsTransitAvailable,
			&room.IsDailyAvailable,
			&room.IsMonthlyAvailable,
			&room.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r *Repository) FindByID(id string) (*Room, error) {
	var room Room

	err := r.DB.QueryRow(`SELECT * FROM rooms WHERE room_id=$1`, id).Scan(
		&room.RoomID,
		&room.Name,
		&room.Location,
		&room.PriceTransit,
		&room.PriceDaily,
		&room.PriceMonthly,
		&room.DepositAmount,
		&room.IsActive,
		&room.IsTransitAvailable,
		&room.IsDailyAvailable,
		&room.IsMonthlyAvailable,
		&room.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *Repository) Update(room *Room) error {
	query := `
	UPDATE rooms SET
		name=$1,
		location=$2,
		price_transit=$3,
		price_daily=$4,
		price_monthly=$5,
		deposit_amount=$6,
		is_active=$7,
		is_transit_available=$8,
		is_daily_available=$9,
		is_monthly_available=$10
	WHERE room_id=$11
	`

	result, err := r.DB.Exec(
		query,
		room.Name,
		room.Location,
		room.PriceTransit,
		room.PriceDaily,
		room.PriceMonthly,
		room.DepositAmount,
		room.IsActive,
		room.IsTransitAvailable,
		room.IsDailyAvailable,
		room.IsMonthlyAvailable,
		room.RoomID,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return utils.NewApiError(404, "room not found")
	}

	return nil
}

func (r *Repository) Delete(id string) error {
	result, err := r.DB.Exec(`DELETE FROM rooms WHERE room_id=$1`, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return utils.NewApiError(404, "room not found")
	}

	return nil
}
