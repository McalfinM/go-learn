package account

import (
	"database/sql"

	"go-api/internal/utils"
)

type RoleRepository struct {
	DB *sql.DB
}

func NewRoleRepository(db *sql.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (r *RoleRepository) Create(role *Role) error {
	return r.DB.QueryRow(
		"INSERT INTO roles(name) VALUES($1) RETURNING role_id",
		role.Name,
	).Scan(&role.RoleID)
}

func (r *RoleRepository) FindAll() ([]Role, error) {
	rows, err := r.DB.Query("SELECT role_id, name FROM roles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []Role

	for rows.Next() {
		var role Role
		rows.Scan(&role.RoleID, &role.Name)
		roles = append(roles, role)
	}

	return roles, nil
}

func (r *RoleRepository) Delete(id int) error {
	res, err := r.DB.Exec("DELETE FROM roles WHERE role_id=$1", id)
	if err != nil {
		return err
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		return utils.NewApiError(404, "role not found")
	}

	return nil
}

func (r *RoleRepository) FindByName(name string) (*Role, error) {
	var role Role

	err := r.DB.QueryRow(
		"SELECT role_id, name FROM roles WHERE name=$1",
		name,
	).Scan(&role.RoleID, &role.Name)

	if err != nil {
		return nil, nil // tidak ditemukan
	}

	return &role, nil
}
