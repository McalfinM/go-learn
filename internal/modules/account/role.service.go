package account

import (
	"go-api/internal/utils"
)

type RoleService struct {
	RoleRepo *RoleRepository
}

func NewRoleService(s *RoleRepository) *RoleService {
	return &RoleService{RoleRepo: s}
}

func (s *RoleService) Create(role *Role) error {
	if role.Name == "" {
		return utils.NewApiError(400, "role name is required")
	}

	// optional: validasi duplikat
	existing, _ := s.RoleRepo.FindByName(role.Name)
	if existing != nil {
		return utils.NewApiError(400, "role already exists")
	}

	return s.RoleRepo.Create(role)
}

func (s *RoleService) GetAll() ([]Role, error) {
	return s.RoleRepo.FindAll()
}

func (s *RoleService) Delete(id int) error {
	if id <= 0 {
		return utils.NewApiError(400, "invalid role id")
	}

	return s.RoleRepo.Delete(id)
}
