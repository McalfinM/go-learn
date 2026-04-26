package room

import (
	"go-api/internal/utils"
)

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateRoom(r *Room) error {

	if r.Name == "" {
		return utils.NewApiError(400, "Nama cant be empty")
	}

	if r.PriceTransit <= 0 || r.PriceDaily <= 0 || r.PriceMonthly <= 0 {
		return utils.NewApiError(400, "price must be greater than 0")
	}

	return s.Repo.Create(r)
}

func (s *Service) UpdateRoom(room *Room) error {
	if room.RoomID == "" {
		return utils.NewApiError(400, "room_id is required")
	}

	return s.Repo.Update(room)
}

func (s *Service) GetRoom() ([]Room, error) {
	return s.Repo.FindAll()
}
