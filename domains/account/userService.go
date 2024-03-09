package account

import "learn/api/databases/model"

type Service interface {
	FindAll() ([]model.User, error)
	FindById(Id int) (model.User, error)
	Create(user RegisterInput) (model.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]model.User, error) {
	return s.repository.FindAll()
}

func (s *service) FindById(Id int) (model.User, error) {
	user, err := s.repository.FindById(Id)

	return user, err
}

func (s *service) Create(request RegisterInput) (model.User, error) {
	entity := model.User{
		Username: request.Username,
		FullName: request.FullName,
		Password: request.Password,
	}
	user, err := s.repository.Create(entity)

	return user, err
}
