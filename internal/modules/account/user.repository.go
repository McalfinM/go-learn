package account

import (
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(email, password string) error {
	_, err := r.DB.Exec(`
		INSERT INTO users (email, password, role_id)
		VALUES ($1, $2, $3)
	`, email, password, 3)

	return err
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
	var user User

	err := r.DB.QueryRow(`
		SELECT u.id, u.user_uuid, u.email, u.password, u.role_id, r.name as role_name
		FROM users u
		JOIN roles r ON r.id = u.role_id
		WHERE u.email = $1
	`, email).Scan(
		&user.ID,
		&user.UserUUID,
		&user.Email,
		&user.Password,
		&user.RoleID,
		&user.RoleName,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByUuid(user_uuid string) (*User, error) {
	var user User

	err := r.DB.QueryRow(`
		SELECT u.id, u.user_uuid, u.email, u.role_id, r.name as role_name
		FROM users u
		JOIN roles r ON r.id = u.role_id
		WHERE u.user_uuid = $1
	`, user_uuid).Scan(
		&user.ID,
		&user.UserUUID,
		&user.Email,
		&user.Password,
		&user.RoleID,
		&user.RoleName,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
